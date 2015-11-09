## simple

A simple console application that shows off the very basics of submarine. The script automatically walks through and logs a typical submarine workflow, which includes the following steps:

1. Creating a new app on submarine.
	* submarine gives the application an API-KEY that can be used to create users
2. "Signing up" a new user
	* when a user signs up for an application that uses submarine, the application has to notify submarine about this.
	* submarine stores a reference to the user and creates a unique, random key that can be used to securely store data. This is returned to the calling application.
3. Storing encrypted data
	* the sample application stores some arbitrary data. Before saving it, however, the data is encrypted with the user's key.
4. Showing the encrypted data
5. Decrypting the data
	* The user's key is used to unlock the data and show it again.

### run it:

`go run samples/simple/test.go`