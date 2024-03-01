# Chat With Code

<div align="center">
  <a href="https://github.com/emilkje/go-openai-toolkit">
    <img src="docs/assets/yelling_at_code.webp" alt="Logo">
  </a>
</div>

## Overview

Chat With Code is yet another command-line application that bridges the gap between coding and conversation. This tool lets developers engage with their codebases using natural language prompts, offering a fresh approach to code exploration and problem-solving.

## Features

- **Interactive Chat Sessions**: Start a dialogue with your codebase to learn about its structure, get summaries of different parts, or even debug issues.
- **Intelligent Context-Aware Responses**: Powered by OpenAI, Chat With Code understands the context of your project, providing meaningful insights and relevant code snippets.
- **Customizable File Inclusion**: Filter the files you want the tool to consider using regular expressions, ensuring focused and relevant chat interactions.
- **Gitignore Awareness**: Exclude files listed in `.gitignore` from the chat context to maintain confidentiality and relevance.
- **Simplicity**: A simple and intuitive interface that requires minimal setup to get started.

## Installation

### Using Go

If you have Go installed (version 1.22 or higher), you can install Chat With Code using the following command:

```sh
go install github.com/emilkje/cwc@latest
```

### Pre-built Binaries

We also provide pre-built binaries for Windows, macOS, and Linux. You can download them from the [releases page](https://github.com/emilkje/cwc/releases) on GitHub.

### Prerequisites

Before you start using Chat With Code, make sure to set the following environment variables:

- `AOAI_API_KEY`: Your OpenAI API key.
- `AOAI_ENDPOINT`: The OpenAI API endpoint.
- `AOAI_API_VERSION`: The OpenAI API version.
- `AOAI_MODEL_DEPLOYMENT`: The deployment of your OpenAI model.

### Getting Started

Once you've installed Chat With Code, you can interact with your code with:

```sh
cwc
```

Feel free to customize the scanning behavior with available flags such as `--include` to specify file patterns 
or `--paths` to define the directories to include in the session. For more information, run:

```sh
cwc --help
```

## Roadmap 

> Note: these items may or may not be implemented in the future.

- [ ] tests
- [ ] support both azure and openai credentials
- [ ] `cwc login` command to set up credentials
- [ ] CI/CD for releases
- [ ] chat using web ui with `cwc web`
- [ ] indexing/search implementation for large codebases
- [ ] tools for dynamic context awareness
- [ ] tools for gathering external documentation

## Contributing

Please file an issue if you encounter any problems or have suggestions for improvement. We welcome contributions in the form of pull requests, bug reports, and feature requests.

## License

Chat With Code is provided under the MIT License. For more details, see the [LICENSE](LICENSE) file.

If you encounter any issues or have suggestions for improvement, please open an issue in the project's [issue tracker](https://github.com/emilkje/chat-with-code/issues).

[banner-photo-url]: ./docs/assets/yelling_at_code.webp