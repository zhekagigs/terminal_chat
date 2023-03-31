# askGPT-3.5 Turbo Chatbot

This is a simple chatbot implementation that uses the GPT-3.5 language model from OpenAI to generate responses to user input. The user can input text messages, and the chatbot will respond with generated text that attempts to continue the conversation. The user can also input code snippets, and the chatbot will respond with generated code that attempts to continue the conversation. The chatbot uses the GPT-3.5 Turbo API, which is a faster version of the GPT-3 API.

This project is built using the Go programming language, also known as Golang. One of the main reasons for choosing Golang for this project is its cross-platform support. Golang binaries can be easily compiled for different operating systems and architectures, making it easy to run the program on any system without the need for any additional dependencies or installations. This makes it a great choice for building command-line tools that need to run on different systems with minimal setup requirements. 

## Prerequisites
go version 1.16 darwin/amd64
OR
ubuntu 20.04 if running pre-compiled binary

Download and install Go from https://go.dev/dl/

Read about cross-compiling Go apps: https://freshman.tech/snippets/go/cross-compile-go-programs/


## Getting Started

To use this chatbot, you will need an OpenAI API key. You can sign up for OpenAI and obtain an API key [here](https://beta.openai.com/signup/).

I just want to run the chatbot from the command line on linux:

```
export OPENAI_API_KEY=your_api_key
./build/askGPT --files=path/to/your/input/mycode.go
```

Clone this repository and navigate to the project directory:

```
git clone https://github.com/your_username/gpt3-chatbot.git
cd gpt3-chatbot
```

Set your OpenAI API key as an environment variable:

```
export OPENAI_API_KEY=your_api_key
```

Then, run the chatbot with the following command:

```
go run main.go --files=path/to/your/input/files.go
```

Alternatively, you can build the chatbot binary and run it from the command line:

```
go build -o askGPT
./askGPT --files=path/to/your/input/files.go
```
Or you can run the chatbot from the command line using pre-build binary:
    
```
./build/askGPT --files=path/to/your/input/files.go
```


where `path/to/your/input/files.go` is a comma-separated list of file paths containing initial system input text or code for the chatbot.

## Usage
On first run it will read the content of the files, passed as arguments.
The chatbot will prompt the user for input and then generate a response using GPT-3.5 Turbo API. 

## Contributing

If you would like to contribute to this project, please open a pull request with your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.