#include <stdio.h>
#include <stdlib.h> // atoi

int getfuel(int mass) {
	int fuel;
	fuel = (mass / 3) - 2; // typical div should handle rounding down

	return fuel;
}

int main(int argc, char **argv) {
	int fuel, mass;
	char *line;
	ssize_t linelen;
	size_t linecap;
	unsigned long long total;
	FILE *fp;

	fp = fopen("./input", "r");
	if (!fp) {
		fprintf(stderr, "Cannot open input file\n");
		return 1;
	}

	total = 0;
	for(linelen = 0, line = NULL; (linelen = getline(&line, &linecap, fp)) > 0; ) {
		mass = atoi(line);
		fuel = getfuel(mass);
		//printf("%d\n", fuel);
		total += fuel;

		linelen = 0;
		line = NULL;
	}

	printf("%llu\n", total);
	return 0;
}
