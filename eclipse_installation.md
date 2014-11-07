1. First download eclipse from the official site and download it. 


2. Move it to /opt directory


3. Make the eclipse easier to access 

    ```
    vim /usr/share/applications/eclipse.desktop
    ```
    
    Add the following lines to `eclipse.desktop`
    
    ```
    [Desktop Entry]
    Name=Eclipse 4
    Type=Application
    Exec=/opt/eclipse/eclipse
    Terminal=false
    Icon=/opt/eclipse/icon.xpm
    Comment=Integed Development Environment
    NoDisplay=false
    Categories=Development;IDE;
    Name[en]=Eclipse
     ```
     
    
