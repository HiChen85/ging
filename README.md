# ging
ging is a tool for create gin web framework development templates

This tool is for the freshmen who want to learn golang and gin web framework, it can easily create a gin dev template for you.
This project is also my first formal open-source project. I hope this can help those who are interested in Go programming language
to quickly build a template that helps them understand how a web application works with Go.

At present, there are still many parts of the project that have not been added to the template, and I will gradually improve it in the future.
If you think this tool is helpful, please give me a star. That is really important to me to maintain this project.

## How to use
Before you start to use this tool. I just imagine that you have installed the golang development environment on your device.
If you don't install it, you need to **download** and **install** it **first**.
Make sure the go command has been input in your **PATH environment**.

### Linux/MacOS
1. For the users of these platforms, you can download the tool from the release.
2. After you download the tool, you can decompress it to somewhere you like. I suggest that you can decompress it to an empty folder and run it there.
3. Let's try it:
   1. `cd ~` to go to you home directory
   2. `mkdir gingTemplates && cd gingTemplates` to create an empty folder
   3. Decompress the executable tool **ging** to  **gingTemplates**.
   4. Then use `./ging create your_project_name` to run the tool
   5. You can use the flag `-m` or `--module` to set you go module name. Default module name is demo

### Windows
For Windows User, you can just download the ging_xxx_.exe file and use the command line window to run it.
I recommend that you to run the tool with powershell as that is closer linux cli.

If you guys find any problems while using this tool, please send your problems on Github Issues.