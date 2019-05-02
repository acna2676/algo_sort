#include <iostream>
#define N 10

using namespace std;

void bubblesort(int *array,int size){
	for(int i = 0; i < size; i++){
		for(int j = i + 1; j < size; j++){
			if(array[i] > array[j]){
				int number = array[i];
				array[i] = array[j];
				array[j] = number;
			}
		}
	}
}

int main(void){
	int array[N];
        for(int i = 0; i < N; i++){
		array[i] = rand();
	}
    
        bubblesort(array,N);

	cout << "結果" << endl;

	for(int i = 0; i < N; i++){
		cout << array[i] << endl;
	}
}