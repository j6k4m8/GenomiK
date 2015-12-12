# To start
Please install `go` using `README-GO.md` first as this will result in you having a cloned version of our repo!
# How to Install Nodejs on UGRAD
Fortunately, node is already installed on all ugrad machines! :-D

# How to Install Meteor on UGRAD
Installing `meteor` is slightly more complicated but overall still easy.
Please note that meteor will atempt to install itself globally (by default in `/usr/local/*`).
As a result, it will request `sudo` access.

In order to install `meteor` locally (which we will _assume you are doing_), deny it `sudo` access by refusing to provide a root account password (keep hitting enter...sorry).

To install, run: _(deny sudo access when requested)_
```
curl https://install.meteor.com | sh
```
This will install `meteor` in `$HOME/.meteor`. You should now add this directory to your `$PATH` by adding this to your `~/.bashrc` file:
```
export PATH=$PATH:$HOME/.meteor
```
Now reload that file by running:
```
source ~/.bashrc
```

You should now have `meteor` installed and in your path.
