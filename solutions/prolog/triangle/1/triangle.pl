triangle(Side1, Side2, Side3, Type) :-
    isTriangle(Side1, Side2, Side3), Side1 = Side2, Side2 = Side3, Type = "equilateral";
    isTriangle(Side1, Side2, Side3), Side1 = Side2, Type = "isosceles";
    isTriangle(Side1, Side2, Side3), Side1 = Side3, Type = "isosceles";
    isTriangle(Side1, Side2, Side3), Side2 = Side3, Type = "isosceles";
    isTriangle(Side1, Side2, Side3), Side1 \== Side2, Side1 \== Side3,
    Side2 \== Side3, Type = "scalene".

isTriangle(A, B, C) :-
    A > 0,
    B > 0,
    C > 0,
    A + B >= C,
    B + C >= A,
    A + C >= B.
