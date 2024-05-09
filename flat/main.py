import os
import zipfile
import random
import string
import subprocess
import requests
from bs4 import BeautifulSoup
import shutil


def download_libraries(output_dir, num_libraries):
    url = "https://mvnrepository.com/"
    response = requests.get(url)
    soup = BeautifulSoup(response.text, "html.parser")
    library_links = []
    for link in soup.find_all("a"):
        href = link.get("href")
        if href and href.startswith("/artifact/") and href.endswith(".jar"):
            library_links.append(href)
    library_links = library_links[:num_libraries]
    for index, link in enumerate(library_links, start=1):
        library_url = url + link
        library_name = link.split("/")[-1]
        print(f"Downloading library {index}/{num_libraries}: {library_name}")
        library_path = os.path.join(output_dir, library_name)
        with open(library_path, "wb") as f:
            f.write(requests.get(library_url).content)


def compile_java_source():
    subprocess.run(["javac", "Main.java"])


def generate_jar_file(
    jar_name, package_name, num_embedded_jars, max_embedded_jar_size_mb, libraries_dir
):
    package_dir = os.path.join(package_name.replace(".", "/"))
    os.makedirs(package_dir, exist_ok=True)

    #    compile_java_source()

    shutil.copyfile("Main.class", os.path.join(package_dir, "Main.class"))

    with zipfile.ZipFile(jar_name, "w") as jar_file:
        main_class_path = os.path.join(package_dir, "Main.class")
        jar_file.write(main_class_path)

        for i in range(num_embedded_jars):
            embedded_jar_name = (
                "".join(random.choices(string.ascii_lowercase + string.digits, k=8))
                + ".jar"
            )
            embedded_jar_size = (
                random.randint(1, max_embedded_jar_size_mb) * 1024 * 1024
            )
            with zipfile.ZipFile(embedded_jar_name, "w") as embedded_jar_file:
                for j in range(embedded_jar_size // (1024 * 1024)):
                    embedded_jar_file.writestr(
                        f"data_{j}.txt",
                        "".join(random.choices(string.ascii_lowercase, k=1024 * 1024)),
                    )
                jar_file.write(
                    embedded_jar_name,
                    f"{package_dir}/embedded_jars/{embedded_jar_name}",
                )
            os.remove(embedded_jar_name)

        for root, dirs, files in os.walk(libraries_dir):
            for file in files:
                if file.endswith(".jar"):
                    jar_file.write(
                        os.path.join(root, file), f"{package_dir}/libraries/{file}"
                    )


if __name__ == "__main__":
    jar_name = "large_jar_with_embedded_jars_and_libraries.jar"
    package_name = "com.example"
    num_embedded_jars = 10
    max_embedded_jar_size_mb = 100
    num_libraries = 100
    libraries_dir = "libraries"

    if not os.path.exists(libraries_dir):
        os.makedirs(libraries_dir)
        download_libraries(libraries_dir, num_libraries)

    generate_jar_file(
        jar_name,
        package_name,
        num_embedded_jars,
        max_embedded_jar_size_mb,
        libraries_dir,
    )
    print(
        f"Generated {jar_name} with {num_embedded_jars} embedded JARs, each up to {max_embedded_jar_size_mb} MB in size, and {num_libraries} imported libraries in package {package_name}."
    )
