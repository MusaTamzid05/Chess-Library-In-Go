import requests

def draw_board():
    res = requests.get("http://localhost:8080/board")
    board = res.json()

    for row in board:
        print(row)


if __name__ == "__main__":
    running = True

    while running:
        prompt = input(">>> ").split()
        choice = prompt[0]

        if choice == "draw":
            draw_board()


        if choice == "exit":
            running = False


