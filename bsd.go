// +build freebsd openbsd netbsd

package gopass

/*
#include <termios.h>
#include <unistd.h>
#include <stdio.h>

int getch(int termDescriptor) {
        int ch;
        struct termios t_old, t_new;

        tcgetattr(termDescriptor, &t_old);
        t_new = t_old;
        t_new.c_lflag &= ~(ICANON | ECHO);
        tcsetattr(termDescriptor, TCSANOW, &t_new);

        ch = getchar();

        tcsetattr(termDescriptor, TCSANOW, &t_old);
        return ch;
}
*/
import "C"

func getch(termDescriptor int) byte {
	return byte(C.getch(C.int(termDescriptor)))
}
