# Guía: Montículos Binarios y Colas de Prioridad

## Ejercicios

### Parte 1: Seguimiento

Dado el siguiente montículo de máximo (del cual sólo se muestran las prioridades de sus elementos):

![Montículo de Máximo](imagenes/heap.drawio.svg)

1. Ejecutar la siguiente secuencia:

    ```go
    Insertar(45)
    Insertar(22)
    Remover()
    Insertar(5)
    ```

    Realice los dibujos del montículo luego de cada momento (tanto el árbol como el arreglo). Puede verificar sus pasos en <https://visualgo.net/en/heap>

2. En un montículo de **mínimo** vacío, ejecutar la secuencia (de la cual sólo conocemos las prioridades de los elementos):

    ```go
    Insertar(10)
    Insertar(12)
    Insertar(1)
    Insertar(14)
    Insertar(6)
    Insertar(5)
    Insertar(8)
    Insertar(15)
    Insertar(3)
    Insertar(9)
    Insertar(7)
    Insertar(4)
    Insertar(11)
    Insertar(13)
    Insertar(2)
    Remover()
    Remover()
    ```

### Parte 2: Código

1. Dado [`Heap`](heap/heap.go), completar la implementación de los métodos en dicho TAD.

    Completar el código para obtener un montículo genérico, cuyos elementos pueden ser cualquier tipo que pueda compararse entre sí, por ejemplo enteros, flotantes, cadenas o estructuras más complejas, como `struct{nombre string, edad int}`, etc. y que soporte montículos de máximos y mínimos.

    > Se recomienda crear métodos adicionales para mejorar la modularidad y reutilización del código. Aunque la interfaz del TAD no debe ser modificada.

2. Crear una [Cola de Prioridad](pqueue/pqueue.go) de [Personas](persona/persona.go) que permita agregar Personas y atenderlas por orden de prioridad y orden de llegada.

    La prioridad se define por la edad de la persona, es decir la persona con mayor edad tiene mayor prioridad.
