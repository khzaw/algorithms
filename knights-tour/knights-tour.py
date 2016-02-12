X = 0
Y = 1

def chess2index(p, boardsize):
    x = boardsize[X] - int(p[1])
    y = ord(p[0]) - ord('a')
    return x, y

def is_legal(pos, board, boardsize):
    return 0 <= (pos[X]) < boardsize[X] and \
           0 <= (pos[Y]) < boardsize[Y] and \
           not board[pos]

def board_string(board, boardsize):
    string = ''
    for x in range(boardsize[X]):
        for y in range(boardsize[Y]):
            string += '%3d' % board[(x, y)]
        string += '\n'
    return string

def next_possible_moves(current, board, boardsize):
    next_moves = [(current[X]+d[X], current[Y]+d[Y]) for d in offsets]
    next_moves = [p for p in next_moves if is_legal(p, board, boardsize)]
    return next_moves

def warnsdorff(current, board, boardsize):
    moves = []
    for move in next_possible_moves(current, board, boardsize):
        temp = board[move]
        board[move] = -1
        moves.append((move, len(next_possible_moves(move, board, boardsize))))
        board[move] = temp
    return sorted(moves, key=lambda x:x[1])

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
    n = int(raw_input('N: '))
    boardsize = (n, n)
    start = chess2index(raw_input('Start: '), boardsize)
    offsets = ((1, 2), (-1, 2), (1, -2), (-1, -2),
               (2, 1), (-2, 1), (2, -1), (-2, -1))
    board = {(x, y):0 for x in range(boardsize[X]) for y in range(boardsize[Y])}
    count = 1
    board[start] = count

    current = start
    while count < len(board):
        p = warnsdorff(current, board, boardsize)
        p = p[0][0]
        count += 1
        board[p] = count
        current = p

    print "Done!"
    print board_string(board, boardsize)


    #  if knight_tour(start, board, count):
        #  print board_string(board, boardsize)
    #  else:
        #  print "No knight's tour"