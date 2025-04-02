Life Weeks Bot
Telegram Bot

A Telegram bot that helps you track how many weeks you've lived and sends weekly reminders about your life progress.

🌟 Features
Calculate and display the total number of weeks you've lived

Weekly notifications every Monday at 7:00 AM with your current "weeks lived" count

Simple and intuitive interface

Persistent storage of user data

Ability to change your birth date

🚀 Getting Started
Prerequisites
Go 1.16 or higher

Telegram Bot Token (get it from @BotFather)

Basic understanding of Telegram bots

Installation
Clone the repository:

bash
Copy
git clone https://github.com/yourusername/life-weeks-bot.git
cd life-weeks-bot
Create a .env file in the root directory with your bot token:

Copy
BOT_TOKEN=your_telegram_bot_token
CONFIG_PATH=configs/main.yml
Install dependencies:

bash
Copy
go mod download
Build and run the bot:

bash
Copy
go run main.go
🛠️ Configuration
The bot uses a YAML configuration file (configs/main.yml) for all messages and settings. You can customize all bot responses there.

Example configuration:

yaml
Copy
messages:
  start: "Welcome! I'll help you track how many weeks you've lived. What's your name?"
  ask_name: "Please tell me your name"
  ask_dob: "Great, % s! Now please send me your date of birth in YYYY-MM-DD format"
  invalid_dob_format: "Sorry, % s, that doesn't look like a valid date. Please use YYYY-MM-DD format"
  invalid_dob: "Sorry, % s, that date doesn't seem right. Please enter a valid date of birth"
  weeks_lived: "% s, you've lived % d weeks so far! I'll remind you every Monday"
  dob_already_set: "Your birth date is already set. Use /change if you want to update it"
  change_dob: "Okay % s, please send me your new date of birth in YYYY-MM-DD format"
  unknown_command: "I don't recognize that command. Use /start to begin"
  start_with: "Let's get started! What's your name?"
🤖 Usage
Start the bot with /start command

Enter your name when prompted

Enter your birth date in YYYY-MM-DD format

The bot will calculate and display your total weeks lived

You'll receive weekly updates every Monday at 7:00 AM

Commands
/start - Begin interaction with the bot

/change - Update your birth date

🧰 Technologies Used
Go Telegram Bot API - Telegram Bot API wrapper for Go

Bolt DB - Embedded key/value database

Cron - Scheduler for weekly notifications

Viper - Configuration management

📂 Project Structure
Copy
.
├── configs/              # Configuration files
├── pkg/
│   ├── storage/          # Database operations
│   ├── telegram/         # Telegram bot handlers
│   └── scheduler/        # Weekly notifications scheduler
├── main.go               # Application entry point
└── README.md             # You are here
