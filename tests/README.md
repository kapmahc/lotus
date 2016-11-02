# Test

## Install ruby
```
git clone https://github.com/rbenv/rbenv.git ~/.rbenv
git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
git clone https://github.com/rbenv/rbenv-vars.git ~/.rbenv/plugins/rbenv-vars
cd ~/.rbenv && src/configure && make -C src
# Zsh note: Modify your ~/.zshrc file instead of ~/.bash_profile
echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> ~/.bash_profile
echo 'eval "$(rbenv init -)"'' >> ~/.bash_profile
```
Then, re-login
```
rbenv install 2.3.1
rbenv local 2.3.1
gem install bundler
bundle install
```

## Run test
```
rspec
```
