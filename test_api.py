import requests

def draw_board():
    res = requests.get("http://localhost:8080/board")
    board = res.json()

    for index, row in enumerate(board):
        print(8 - index, row)

    print("", end="  ")
    print("a b c d e f g h".split())


def show_valid_moves(from_):
    res = requests.post("http://localhost:8080/valid_moves", json={"cell" : from_})
    print(res.json())



if __name__ == "__main__":
    running = True

    while running:
        prompt = input(">>> ").split()
        choice = prompt[0]

        if choice == "draw":
            draw_board()

        elif choice == "valid":
            show_valid_moves(from_= prompt[1])


        if choice == "exit":
            running = False


