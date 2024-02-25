# KYC_Proof

KYC_Proof is a Go program designed to generate a Merkle proof for a specific target hash within a Merkle tree. This tool reads a Merkle tree and a target hash from JSON files, calculates the proof, and outputs the proof to another JSON file.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.15 or higher recommended)
- Git (optional, for cloning the repository)

### Installation

First, clone the repository if available or simply ensure you have the Go file (`main.go`) ready in your workspace:

```sh
git clone https://your-repository-url/KYC_Proof.git
cd KYC_Proof
```

Make sure your project directory structure is as follows:

```
/KYC_Proof
│
├── cmd
│   └── main.go           # Main program file
├── data
│   ├── merkletree.json   # Input Merkle tree file
│   └── targethash.json   # Input target hash file
└── README.md
```

### Setting Up Your Data

1. **Merkle Tree (`merkletree.json`)**: This file should contain your Merkle tree structure. Here is an example format:

```json
[
  ["leaf1Hash", "leaf2Hash"],
  ["parent1Hash"],
  ["rootHash"]
]
```

2. **Target Hash (`targethash.json`)**: This file should specify the hash for which you wish to generate a proof. Example:

```json
{
  "targetHash": "leaf1Hash"
}
```

Replace `"leaf1Hash"` with the actual hash present in your `merkletree.json`.

### Running the Program

Navigate to the `cmd` directory within the KYC_Proof project, then execute the following command:

```sh
go run main.go
```

This command reads the `merkletree.json` and `targethash.json` from the `data` directory, calculates the Merkle proof for the given target hash, and outputs the proof to `data/proof.json`.

### Output

After successful execution, you'll find the generated proof in `data/proof.json`, structured as follows:

```json
{
  "targetLeafHash": "leaf1Hash",
  "proof": ["siblingHash", "parentSiblingHash"]
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Contribute

If you would like to contribute please contact us. 
