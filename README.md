# jitter-rng-go

Implementación educativa de un Generador de Números Aleatorios Verdaderos (TRNG)
basado en **jitter de CPU**, inspirado en:
- [Jitter RNG](https://www.chronox.de/jent/) de Stephan Müller
- `crypto/internal/fips140/entropy` de Go

## Características

- ✅ Fuentes de entropía físicas (sin red/disco): jitter de CPU + ruido de memoria
- ✅ Acondicionamiento determinista con SHA-384 (SP 800-90B)
- ✅ Pool de entropía con **forward secrecy**
- ✅ Health tests continuos (Repetition Count + Adaptive Proportion)
- ✅ API simple y segura
- ✅ Cross-platform: Linux (clock_gettime / TSC) y Windows (QPC / TSC)

## Conceptos clave

### Jitter de CPU
El tiempo de ejecución de secuencias deterministas de instrucciones (bucles,
accesos a memoria) varía debido a efectos no deterministas del hardware:
cachés, pipelines, predicción de saltos, variación de voltaje térmico.

### Scratch Buffer (Memory Access Noise)
Un buffer de 32 MB fuerza accesos a RAM (no caben en L3), introduciendo
latencia variable según el estado del sistema de memoria.

### Forward Secrecy
Tras cada extracción, el estado del pool se rehashea, de modo que comprometer
una salida no compromete generaciones futuras.

## Uso

```go
rng := jitterrng.New()
b, _ := rng.GetRandomBytes(32)
n, _ := rng.GetRandomInt(1, 100)
f, _ := rng.GetRandomFloat64()
