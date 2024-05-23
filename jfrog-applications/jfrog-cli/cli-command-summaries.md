# CLI Command Summaries

## Overview

The command summaries feature allows the recording of JFrog CLI command outputs into the local file system. This can be used to display a summary of the executed CLI command.

Each command execution that has implemented this feature can save data files into the file system. These files will then be used to create a summary markdown.

## How to use?

To use the command summaries, you'll need to set the `JFROG_CLI_COMMAND_SUMMARY_OUTPUT_DIR` environment variable.

After that, supported commands will start saving data files to this directory, along with markdown files.

### ⚠️ Notice: Delete your files after multiple executions

The CLI doesn't delete the files itself as they need to persist after more than one execution, so note that it is your responsibility to clear after you in your pipelines.

## How to implement? 👩‍💻

If you'd like to implement your own summary, you can follow these steps:

1. Implement the CommandSummaryInterface
2. Record data during runtime

#### Implement the CommandSummaryInterface
 ```go
type CommandSummaryInterface interface {
	GenerateMarkdownFromFiles(dataFilePaths []string) (finalMarkdown string, err error)
}
 ```
  
#### Record data during runtime
```go
    // Initialize your implementation
	myNewCommandSummary, err := commandsummary.New(&MyCommandStruct{}, "myNewCommandSummary")
	if err != nil {
		return
	}
	// Record
	return myNewCommandSummary.Record(data)
 ```

The function `GenerateMarkdownFromFiles` needs to handle multiple data files, which are the results of previous command executions, and produce one markdown string content.
As each CLI command has its own context, that's the reason each time we need to reproduce the entire markdown with the new added results.

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

 

## How does it work? 🚧

Each command that implements the `CommandSummaryInterface` will have its own subdirectory inside the `JFROG_CLI_COMMAND_SUMMARY_OUTPUT_DIR/JFROG_COMMAND_SUMMARY` directory.

Each subdirectory will hold data files, each representing a command recording, and a markdown file which has been generated from the entire data files.


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

The user-implemented function will take all of the data files inside its subdirectory and will produce a markdown string.

#### ❓ Why do we need to generate markdown from all the files every time?

The reason for this is that the CLI commands execution does not share a common context, and that is why the markdown file should be generated every time, assuming it is the last command that will run. As we can't know for sure when the last command will be executed, every time we call record, we also have to generate the markdown from all the existing files in the directory.