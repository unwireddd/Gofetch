#!/usr/bin/env bash

# Set a trap for SIGINT (CTRL+C) and any errors
trap ERR SIGINT

# Begin the installation process
echo "Starting installation..."

# Step 1: Create directory
if [ -d "/etc/Gofetch" ]; then
    echo "Directory /etc/Gofetch already exists. Skipping creation."
else
    if sudo mkdir /etc/Gofetch; then
        echo "Directory /etc/Gofetch created successfully."
    else
        echo "Failed to create directory /etc/Gofetch."
        exit 1
    fi
fi

# Step 2: Copy the Ascii directory
if [ -d "/etc/Gofetch/Ascii" ]; then
    echo "Directory /etc/Gofetch already exists. Skipping creation."
else
    if sudo cp -r Ascii /etc/Gofetch; then
        echo "Ascii directory copied successfully."
    else
        echo "Failed to create directory /etc/Gofetch/Ascii."
        exit 1
    fi
fi

# Step 3: Make the gofetch script executable if not already executable
if [ -x "gofetch" ]; then
    echo "gofetch application is already executable. Skipping chmod."
else
    if sudo chmod +x gofetch; then
        echo "gofetch application made executable."
    else
        echo "Failed to make gofetch application executable."
        exit 1
    fi
fi

# Step 4: Copy the gofetch script to /usr/local/bin if not already present
if [ -f "/usr/local/bin/gofetch" ]; then
    echo "gofetch script already exists in /usr/local/bin. Skipping copy."
else
    if sudo cp gofetch /usr/local/bin; then
        echo "gofetch script copied to /usr/local/bin."
    else
        echo "Failed to copy gofetch script to /usr/local/bin."
        exit 1
    fi
fi

# Installation complete
echo "Your installation is done!"
trap - ERR SIGINT
