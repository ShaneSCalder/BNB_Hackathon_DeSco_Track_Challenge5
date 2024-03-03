# Video Merkle Tree Generator

This Go application demonstrates the construction of a Merkle tree from a video file. It chunks the video into fixed-size pieces, builds a Merkle tree from these chunks, calculates the Merkle root, and outputs both the Merkle root and a visual representation of the entire Merkle tree.

## Features

- **Video Chunking**: Splits a video file into manageable chunks.
- **Merkle Tree Construction**: Builds a Merkle tree from the video chunks.
- **Merkle Root Output**: Saves the Merkle root to `data/merkleroot.txt`.
- **Tree Representation**: Generates a textual representation of the Merkle tree in `data/merkletree.txt`.

## Getting Started

These instructions will help you set up and run the project on your local machine.

### Prerequisites

Ensure you have [Go](https://golang.org/dl/) installed on your system. This project is compatible with Go 1.15 and newer versions.

### Installation

1. Clone this repository to your local machine:

2. Navigate to the project directory:

### Usage

To run the program, execute the following command from the root of the project directory:

bash
go run cmd/main.go


## Usage

The program expects a video file named `TestVideo.mp4` in the `data` directory by default.
Ensure this file exists, or modify the `videoPath` variable in `cmd/main.go` to point to the location of your video file.

## Outputs

After running, the program generates two files in the `data` directory:

- `merkleroot.txt`: Contains the Merkle root of the constructed tree.
- `merkletree.txt`: Provides a visual representation of the Merkle tree.

## Contributing

Please feel free to contribute to this project.

## License

This project is licensed under the MIT License.

## Compiling for Use in Web Applications

After adjusting the Go program according to your needs, you can compile it to a binary executable for various platforms.
This compiled version can then be integrated into web applications, enabling server-side processing or even client-side use via WebAssembly (depending on the use case and compilation target).

markdown
### Compiling the Go Program

To compile the program for your current platform, run:

bash
go build -o merkletree cmd/main.go


This command will create an executable file named `merkletree` (or `merkletree.exe` on Windows) that can be used directly from the command line, integrated into Visual Studio Code tasks, or called from other programs and web applications.

### Integration with Other Technologies

Once compiled into an executable, the Go program can be utilized within a variety of programming and scripting environments. This cross-language compatibility opens up numerous possibilities for integrating advanced Go functionalities into broader applications and systems.

- **In Command Line Scripts**: Directly execute the binary from batch scripts (`.bat`) on Windows or shell scripts (`.sh`) on Unix-like systems.
- **Within Visual Studio Code**: Use tasks or the integrated terminal in Visual Studio Code to run the executable as part of your development workflow.
- **With Node.js**: Use the `child_process` module to spawn the executable from a Node.js application, allowing for integration into web servers or backend services.
- **In Python**: Utilize the `subprocess` module to call the executable for processing video files, useful in data analysis or media processing applications.
- **From Web Applications**: For server-side operations in web applications, execute the binary for tasks like video processing or generating Merkle trees for uploaded files.

### Example: Calling from Node.js

Here's a quick example of how you might call the compiled executable from a Node.js application using the `child_process` module:

```javascript
const { exec } = require('child_process');

exec('./merkletree data/TestVideo.mp4', (error, stdout, stderr) => {
    if (error) {
        console.error(`Execution error: ${error}`);
        return;
    }
    console.log(`Merkle Root: ${stdout}`);
});
```

This integration flexibility makes the compiled Go program a versatile tool that can be deployed in a wide range of application scenarios, enhancing both development efficiency and application capabilities.

# Example output data 
two JSON files have been included but program can be modified to write other formats. 
Other formats can be created JSON, XML, Binary Format, CSV, YAML, etc. 

## Merkle Tree Complexity and Future Directions

The construction and management of the Merkle tree for video data encapsulates a unique set of challenges and complexities. Video data, inherently large and requiring nuanced consideration for how it's segmented (chunked) for inclusion in the tree, significantly influences the overall architecture and performance of the Merkle tree.

### Challenges with Video Data

- **Chunking Size**: Determining the optimal chunk size for video data is critical. It directly impacts the depth and breadth of the Merkle tree, affecting computational efficiency for both tree construction and the generation/verification of Merkle proofs. Smaller chunks increase the tree's depth, potentially complicating proof generation, whereas larger chunks could lead to inefficiencies in data management and verification.
  
- **Data Type Complexity**: Video content, characterized by its substantial size, sequential playback requirement, and variable bitrates, presents unique challenges. These include ensuring data integrity across chunks, efficiently accessing specific segments of video data, and accommodating variable sizes and bitrates without compromising the Merkle tree's integrity or performance.

### Towards Simplification

Acknowledging the inherent complexity introduced by the chunking of video data and its impact on the Merkle tree, we are committed to simplifying this process. Our ongoing efforts and future work aim to:

- **Optimize Chunking Strategies**: We're exploring more efficient chunking strategies to reduce the Merkle tree's depth, thereby decreasing the computational overhead associated with tree construction and proof operations.

- **Enhance Data Representation**: Investigating alternative methods for representing video data within the tree to streamline proof generation and verification processes, making them more efficient and less resource-intensive.

- **Improve Computational Efficiency**: By refining our approach to constructing and managing the Merkle tree, we aim to enhance the overall efficiency, making it more suitable for large-scale video data applications.

Our goal is to balance the need for data integrity and verification with the practical considerations of handling large video files, ensuring that our Merkle tree implementation remains robust yet efficient. We welcome contributions and insights from the community as we navigate these challenges and work towards more streamlined solutions.
