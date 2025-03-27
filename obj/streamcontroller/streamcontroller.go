package streamcontroller

import (
	"fmt"
	"sync"

	"github.com/fengdotdev/golibs-streams/obj/stream"
)

type StreamControllerInterface[T any] interface {
	Stream() stream.Stream[T]
	Sink()
	Add(item T)
}

type StreamController[T any] struct {
	mu          sync.Mutex
	stream      stream.Stream[T]
	closed      bool
	subscribers []chan interface{}
}

func NewStreamController[T any]() StreamController[T] {
	return StreamController[T]{stream: stream.NewStreamEmpty[T]()}
}

func (sc *StreamController[T]) Stream() *stream.Stream[T] {
	return &sc.stream
}

func (sc *StreamController[T]) Add(item T) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if sc.closed {
		return // No hacer nada si está cerrado
	}

	// Enviar el dato a todos los suscriptores
	for _, sub := range sc.subscribers {
		select {
		case sub <- item: // Enviar de forma no bloqueante si el canal está listo
		default:
			// Si el suscriptor no está listo, podrías manejarlo (ignorar, encolar, etc.)
			fmt.Println("Suscriptor no listo, dato descartado")
		}
	}
}

func (sc *StreamController[T]) Close() {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if !sc.closed {
		sc.closed = true
		for _, sub := range sc.subscribers {
			close(sub)
		}
		sc.subscribers = nil // Limpiar suscriptores
	}
}


func (sc *StreamController[T]) Subscribe() <-chan interface{} {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if sc.closed {
		ch := make(chan interface{})
		close(ch)
		return ch
	}

	// Crear un canal para el nuevo suscriptor
	ch := make(chan interface{}, 1) // Buffer opcional para evitar bloqueos
	sc.subscribers = append(sc.subscribers, ch)
	return ch
}