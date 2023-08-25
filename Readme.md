
# GoPass
Welcome to the CLI Password Manager project! This is a simple command-line application written in Go that helps you manage your passwords securely. It uses a MongoDB database to store your password data. This README provides an overview of the project and how to use it.

## Features

- Add, retrieve, update, and delete passwords for different platforms.
- Generate strong passwords for new accounts.
- Understand the password policy guidelines for better security.
- A user-friendly command-line interface.
- Securely store your password data in a MongoDB database.

## Installation

1. Clone this repository to your local machine.
   ```
   git clone https://github.com/speedhs/gopass.git
   cd gopass
   ```

2. Install the necessary dependencies using Go Modules.
   ```
   go mod tidy
   ```

3. Create a `.env` file in the root directory and set your MongoDB URI as follows:
   ```
   MONGODB_URI=mongodb://username:password@localhost:27017/mydb
   ```

4. Build and run the application.
   ```
   go run main.go
   ```

## Usage

- Launch the application using the command provided in the Installation section.
- Use the displayed menu to interact with the password manager.
- Choose an option by entering the corresponding number.
- Follow the prompts to add, retrieve, update, or delete passwords.

## Contribution

Contributions to this project are welcome! If you have ideas for improvements or new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

This project was inspired by the need for a simple and secure way to manage passwords from the command line.

## Disclaimer

This project is still under development. Use at your own risk. The authors are not responsible for any security breaches or data loss.

## Contact

For any questions or suggestions, please feel free to reach out to me on [Twitter](https://twitter.com/dizzyspeeed).

Enjoy secure password management with GoPass!


---
