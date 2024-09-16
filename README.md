This project uses environment variables to protect sensitive data such as email-address and the app password
You will need to create the .env file in the root of the project directory in the following structure:
1. Create a .env file
2. Put the following data in the .env file
EMAIL_ADDR="your-emailaddress@gmail.com"
EMAIL_PASS="your-app-pass"
Make sure to replace your-emailaddress@gmail.com with your actual Gmail address and your-app-password with the Gmail app password you generate (not your regular Gmail account password)
3.These .env will be autmatically load in your project, ensuring that you credentails remain secure and not hard-coded into the source code
