import os  # For file operations
import time  # For sleeping

for i in range(1):  # Run for 1 iterations
    try:
        with open('./instructions.md', 'r') as file:
            content = file.read().strip()  # Read and strip whitespace
        if not content:  # Check if the file is empty
            time.sleep(120)  # Sleep for 2 minutes (120 seconds)
        else:
            print(f'INSTRUCTIONS: {content}')
    except Exception as e:
        print(f'Error in iteration {i+1}: {e}')
