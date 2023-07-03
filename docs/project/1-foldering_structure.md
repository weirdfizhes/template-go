# {Project Name} - Project Foldering

<img src="../assets/structure.jpg" alt="Golang Logo" style="width: 500px; display: block; margin-left: auto; margin-right: auto; width: 50%; margin-bottom: 30px">

In this code, we define the app into 5 (five) different folders. There are: `cmd`, `config`, `docs`, `src`, and `tool`. Each folder will be explained below.

1. Folder `/cmd`\
This folder contain the main file to run the project, can be a .go file, or .sh file (after build the project to bash script file).

2. Folder `/config`\
This folder contain the configuration file such as database configuration, etc. For example, the database config have a configuration to connecting our app to our database.

3. Folder `/docs`\
This folder contains all the documentation for our app. You can store the assets media, postman json files, deployment configuration, etc.

4. Folder `/src`\
This folder contains our main code to build some logical scheme to run our application. As example, authorization application, feature application etc will be in this folder.

5. Folder `/tool`\
This folder contains all the tool to help the code running the http server, or provide constant variable, hashing tool, request validation etc.