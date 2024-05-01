# Quad Function Checker

This project contains a program called `quadchecker` that takes a string as an argument and displays the name of the matching quad function and its dimensions. If the argument does not match any known quad functions, the program outputs `Not a quad function`.

The project also includes multiple quad functions (`quadA`, `quadB`, `quadC`, `quadD`, `quadE`) that you can use to test the `quadchecker` program.

## Instructions

- **Quad Function Checker**: The `quadchecker` program takes a string argument and displays the name of the matching quad function and its dimensions. If there is more than one match, the program will display all matching functions alphabetically, separated by `||`.

- **Not a Quad Function**: If the provided argument does not match any known quad function, the program outputs `Not a quad function`.

- **Newline Characters**: All answers must end with a newline character (`\n`).

## Usage

### 1. Build the Files

First, you need to build all the necessary files by executing the `build.sh` script:

```shell
./build.sh
```

This script will compile all the quad function Go files (`quadA.go`, `quadB.go`, `quadC.go`, `quadD.go`, `quadE.go`) and the `quadchecker.go` file. If the build is successful, the resulting executable files will be moved to the main directory.

### 2. Execute the Quad Function Checker

Once the build is complete, you can use the following commands to execute the `quadchecker` program with different arguments:

- **Example 1**:

    ```shell
    ./quadA 3 3 | go run .
    ```

    This will check for the matching quad function for the argument "3 3" and display the appropriate output.

- **Example 2**:

    ```shell
    ./quadC 1 1 | go run .
    ```

    This will check for the matching quad functions for the argument "1 1" and display the appropriate output.

- **Example 3**:

    ```shell
    ./quadE 1 2 | go run .
    ```

    This will check for the matching quad functions for the argument "1 2" and display the appropriate output.

- **Example 4**:

    ```shell
    echo 0 0 | go run .
    ```

    This will check for the matching quad functions for the argument "0 0" and display whether it is a quad function or not.

## Notes

- Ensure that you have the necessary permissions to execute the scripts and files (e.g., `chmod +x build.sh`).
- Make sure to compile the Go files first using `build.sh` before executing the `quadchecker` program.

## Author

Created by Hamza Maach.