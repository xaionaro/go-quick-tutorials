all: 00_hello_world/run 01_assignments/run 02_arrays_and_maps/run 03_loops/run n0_reflect/run

00_hello_world/run:
	make -C 00_hello_world

01_assignments/run:
	make -C 01_assignments

02_arrays_and_maps/run:
	make -C 02_arrays_and_maps

03_loops/run:
	make -C 03_loops

n0_reflect/run:
	make -C n0_reflect

clean:
	make -C 00_hello_world clean
	make -C 01_assignments clean
	make -C 02_arrays_and_maps clean
	make -C 03_loops clean
	make -C n0_reflect clean

