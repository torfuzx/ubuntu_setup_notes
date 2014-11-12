- 1. Add the following line to /etc/apt/sources.list
  
  sudo sh -c "echo deb http://repository.spotify.com stable non-free >> /etc/apt/sources.list"

- 2. Add the public key in case you want to verify the downloaded package

  sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 94558F59

3. Run apt-get update

  sudo apt-get update

4. Install spotify!
  
  sudo apt-get install spotify-client
