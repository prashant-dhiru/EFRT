
# EFRT - JIRA Effort Logger CLI

## Overview

EFRT (Effort Logger for JIRA) is a command-line interface (CLI) application designed to simplify the process of logging effort in JIRA. This tool streamlines the effort tracking workflow, allowing users to efficiently update their work logs directly from the terminal.

## Features

- **Effort Logging:** Log your work effortlessly by providing key details such as JIRA issue key, time spent, and optional comments.
- **Interactive Mode:** Use the interactive mode for a user-friendly experience, guiding you through the process of logging effort step by step.

<!-- TODO windowss and linux  executable  -->
## Installation

1. Ensure you have Python installed on your machine. (EFRT requires Python 3.6 or above)
2. Clone the EFRT repository:

    ```bash
    git clone https://github.com/your-username/efrt.git
    ```

3. Navigate to the EFRT directory:

    ```bash
    cd efrt
    ```

4. Install the required dependencies:

    ```bash
    pip install -r requirements.txt
    ```

5. Make the script executable:

    ```bash
    chmod +x efrt.py
    ```

6. Run EFRT:

    ```bash
    ./efrt.py
    ```

## Usage

### Logging Effort

1. **Interactive Mode:**
    ```bash
    ./efrt log
    ```

    Follow the on-screen prompts to log effort interactively.
    #### Log Flags
    - **-c, --comment:** Interactively add comment 
    - **-o, --old:** Interactively add worklog for pervious day


### Global Flag
- **-h, --help:** Display help information

## Configuration

1. Display your JIRA config using the `config` command:

    ```bash
    ./efrt.py config
    ```

    Follow the instructions to provide your JIRA base URL, username, and API token.

2. (Optional) Set up your JIRA config using the `--set` flag such as the JIRA server, JIRA Personal Access Token, etc :

    ```bash
    ./efrt.py config --set=KEY:VALUE
    ```

<!-- ## Example

### Interactive Effort Logging

```bash
./efrt.py log
```

Follow the prompts to enter the required information for logging effort.

### Batch Effort Logging

```bash
./efrt.py batch --file path/to/effort_logs.csv
```

Ensure your CSV file follows the required format for batch logging.

## Contributing

If you would like to contribute to EFRT, please follow the [contribution guidelines](CONTRIBUTING.md). -->

## License

EFRT is licensed under the Apache License Version 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

Special thanks to the [JIRA API](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/) for making this integration possible.

---

**Note:** EFRT is not officially affiliated with or endorsed by JIRA or Atlassian.
