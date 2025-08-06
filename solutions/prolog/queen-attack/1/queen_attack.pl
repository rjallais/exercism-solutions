% Please visit https://exercism.io/tracks/prolog/installation
% for instructions on setting up prolog.
% Visit https://exercism.io/tracks/prolog/tests
% for help running the tests for prolog exercises.

% Replace the goal below with
% your implementation.

%! create(+DimTuple)
%
% The create/1 predicate succeeds if the DimTuple contains valid chessboard 
% dimensions, e.g. (0,0) or (2,4).
create((DimX, DimY)) :-
	DimX > -1,
	DimX < 8,
	DimY > -1,
	DimY < 8.

%! attack(+FromTuple, +ToTuple)
%
% The attack/2 predicate succeeds if a queen positioned on ToTuple is 
% vulnerable to an attack by another queen positioned on FromTuple.
attack((FromX, _), (ToX, _)):-
	FromX = ToX.

attack((_, FromY), (_, ToY)):-
	FromY = ToY.

attack((FromX, FromY), (ToX, ToY)):-
	X is abs(ToX - FromX), Y is abs(ToY - FromY), X = Y.
