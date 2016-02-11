X = 0
Y = 1

def chess2index(p, boardsize):
    x = boardsize[X] - int(p[1])
    y = ord(p[0]) - ord('a')
    return x, y

def is_legal(pos, board, boardsize):
    return 0 <= (pos[X]) < boardsize[X] and \
           0 <= (pos[Y]) < boardsize[Y] and \
           board[pos] == 0

def board_string(board, boardsize):
    string = ''
    for x in range(boardsize[X]):
        for y in range(boardsize[Y]):
            string += '%3d' % board[(x, y)]
        string += '\n'
    return string

def knight_tour(currentpos, board, count):
    if count >= boardsize[X] * boardsize[Y]:
        return True

    for d in offsets:
        newpos = currentpos[X]+d[X], currentpos[Y]+d[Y]
        if is_legal(newpos, board, boardsize):

            count += 1
            board[newpos] = count

            if knight_tour(newpos, board, count):
                return True
            else:
                board[newpos] = 0
                count -= 1
    return False


if __name__ == '__main__':
    n = int(raw_input('n: '))
    boardsize = (n, n)
    start = chess2index(raw_input('Start: '), boardsize)
    offsets = ((1, 2), (-1, 2), (1, -2), (-1, -2),
            (2, 1), (-2, 1), (2, -1), (-2, -1))
    board = {(x, y):0 for x in range(boardsize[X]) for y in range(boardsize[Y])}
    count = 1
    board[start] = count
    if knight_tour(start, board, count):
        print board_string(board, boardsize)
    else:
        print "No knight's tour"