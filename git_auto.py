import subprocess

def run_command(command):
    """Runs a shell command and prints its output."""
    result = subprocess.run(command, shell=True, capture_output=True, text=True)
    if result.returncode != 0:
        print(f"Error: {result.stderr}")
        exit(result.returncode)
    print(result.stdout)

def git_push(directory, commit_message):
    """Automates the git add, commit, and push process."""
    try:
        # Change to the target directory
        subprocess.run(["cd", directory], check=True)
        
        # Run git commands
        run_command("git add .")
        run_command(f"git commit -m \"{commit_message}\"")
        run_command("git push")
    except subprocess.CalledProcessError as e:
        print(f"Command failed: {e}")

if __name__ == "__main__":
    import os

    # Replace with your directory and commit message
    directory = input("Enter the path to your local Git repository: ")
    commit_message = input("Enter your commit message: ")

    if not os.path.isdir(directory):
        print(f"Error: The directory '{directory}' does not exist.")
        exit(1)

    git_push(directory, commit_message)
