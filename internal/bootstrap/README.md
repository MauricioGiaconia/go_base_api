# Bootstrap

Paquete encargado de la **inicialización y ensamblaje** de la aplicación (Composition Root).

## Responsabilidades

- Crear instancias concretas de todos los adapters outbound (MySQL, JWT, etc.).
- Crear instancias de los application services inyectando los adapters.
- Crear instancias de los handlers inbound inyectando los services.
- Exponer un `Container` con todas las dependencias listas para usar.

## Rol en arquitectura hexagonal

Este paquete es el **único lugar** que conoce las implementaciones concretas de todos los adapters. Es donde se "cablea" la aplicación, conectando:

```bash
Adapters Outbound (MySQL, JWT) 
        ↓ implementan
    Output Ports (interfaces)
        ↓ se inyectan en
Application Services (casos de uso)
        ↓ implementan
    Input Ports (interfaces)
        ↓ se inyectan en
Adapters Inbound (Handlers HTTP)
```

## Reglas

- Puede importar de **todas las capas** (es la excepción a la regla de dependencias).
- NO debe contener lógica de negocio.
- Solo crea instancias y las conecta.
- Si agregás un nuevo módulo (ej: `InvoiceHandler`), agregá su cableado acá.

## Nomenclatura

- Archivo: `container.go`
- Struct: `Container`
- Constructor: `NewContainer()`

## Ejemplo de uso

```go
// main.go
func main() {
    c := bootstrap.NewContainer()
    
    r := gin.Default()
    router.SetupRoutes(r, c.AuthHandler)
    r.Run(":8080")
}
```

## Extendiendo el container

Cuando agregues nuevos módulos:

```go
type Container struct {
    DB             *sql.DB
    AuthHandler    *handler.AuthHandler
    InvoiceHandler *handler.InvoiceHandler  // nuevo
}

func NewContainer() *Container {
    // ... infraestructura existente ...
    
    // Nuevo módulo
    invoiceRepo := mysql.NewInvoiceRepository(db)
    invoiceService := service.NewInvoiceService(invoiceRepo)
    invoiceHandler := handler.NewInvoiceHandler(invoiceService)

    return &Container{
        DB:             db,
        AuthHandler:    authHandler,
        InvoiceHandler: invoiceHandler,  // nuevo
    }
}
```
