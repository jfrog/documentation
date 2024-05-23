# CLI Command Summaries

## Overview

The command summaries feature enables the recording of JFrog CLI command outputs into the local file system.
This functionality can be used to generate a summary of the executed CLI command.

Each command execution that incorporates this feature can save data files into the file system.
These files are then used to create an aggregated summary in Markdown format.

Saving data to the filesystem is essential because CLI command executes in separate contexts.
Consequently, each command that records new data should also incorporate any existing data into the aggregated markdown.
This is required because the CLI cannot determine when a command will be the last one executed in a sequence of commands.

## Notes for Developers 🖥️👩🏻‍💻⚙️

To use the command summaries, you'll need to set the `JFROG_CLI_COMMAND_SUMMARY_OUTPUT_DIR` environment variable.
This variable designates the directory where the data files and markdown files will be stored.

### ⚠️ Notice: Delete Your Files After Multiple Executions

The CLI does not delete the files itself as they need to persist after more than one execution.
Therefore, it is your responsibility to clean up your pipelines.

## How to Implement? 👩‍💻

If you wish to implement your own summary, follow these steps:

1. Implement the CommandSummaryInterface
2. Record data during runtime

#### Implement the CommandSummaryInterface
 ```go
type CommandSummaryInterface interface {
	GenerateMarkdownFromFiles(dataFilePaths []string) (finalMarkdown string, err error)
}
 ```

#### Record Data During Runtime
```go
    // Initialize your implementation
	myNewCommandSummary, err := commandsummary.New(&MyCommandStruct{}, "myNewCommandSummary")
	if err != nil {
		return
	}
	// Record
	return myNewCommandSummary.Record(data)
 ```

The `GenerateMarkdownFromFiles` function needs to process multiple data files, which are the results of previous command executions, and generate a single markdown string content. As each CLI command has its own context, we need to regenerate the entire markdown with the newly added results each time.

### Example Implementation ✍️


```go
// Step 1. Implement the CommandSummaryInterface
type CommandStruct struct{}

type singleRecordedObject struct {
	Name string
}

func (cs *CommandStruct) GenerateMarkdownFromFiles(dataFilePaths []string) (markdown string, err error) {
	// Aggregate all the results into a slice
	var recordedObjects []*singleRecordedObject
	for _, path := range dataFilePaths {
		var singleObject singleRecordedObject
		if err = commandsummary.UnmarshalFromFilePath(path, &singleObject); err != nil {
			return
		}
		recordedObjects = append(recordedObjects, &singleObject)
	}

	// Create markdown
	markdown = results.String()
	return
}

// Step 2. Record data during runtime
func recordCommandSummary(data any) (err error) {
	if !commandsummary.ShouldRecordSummary() {
		return
	}

	commandSummaryImplementation, err := commandsummary.New(&CommandStruct{}, "CommandName")
	if err != nil {
		return
	}

	return commandSummaryImplementation.Record(data)
}

 ```



## How Does It Work? 🚧

Each command that implements the `CommandSummaryInterface` will have its own subdirectory inside the `JFROG_CLI_COMMAND_SUMMARY_OUTPUT_DIR/JFROG_COMMAND_SUMMARY` directory.

Each subdirectory will contain data files, each representing a command recording, and a markdown file which has been generated from all the data files.


```
JFROG_CLI_COMMAND_SUMMARY_OUTPUT_DIR/JFROG_COMMAND_SUMMARY  
│
└─── Command1
│       datafile1.txt
│       datafile2.txt
│       markdown.txt
│   
└─── Command2
        datafile1.txt
        datafile2.txt
        markdown.txt
```

The user-implemented function will process all of the data files inside its subdirectory and will generate a markdown string.

#### ❓ Why Do We Need to Generate Markdown from All the Files Every Time?

The reason for this is that the CLI commands execution does not share a common context. 
Therefore, the markdown file should be regenerated every time, assuming it is the last command that will run.
Since we cannot determine when the last command will be executed, every time we call record,
we also have to generate the markdown from all the existing files in the directory.