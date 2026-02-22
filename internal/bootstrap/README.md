# Bootstrap

Paquete encargado de la **inicialización y ensamblaje** de la aplicación (Composition Root).

## Responsabilidades

- Crear la infraestructura compartida (DB, JWT, etc.).
- Crear middlewares que requieren inyección de dependencias.
- Delegar el ensamblaje de cada dominio a su `*_module.go` correspondiente.
- Ensamblar el slice `[]router.RouteRegistrar` con los módulos de dominio.
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
        ↓ se agregan como
    []router.RouteRegistrar
```

## Reglas

- Puede importar de **todas las capas** (es la excepción a la regla de dependencias).
- NO debe contener lógica de negocio.
- `container.go` solo crea infraestructura compartida y llama a los `*_module.go`.
- Cada dominio tiene su propio archivo `<dominio>_module.go` que encapsula su cableado completo.
- Si agregás un nuevo dominio (ej: `Invoice`), creá `invoice_module.go` y sumá `newInvoiceModule(...)` al slice `Registrars` en `container.go`.

## Nomenclatura

- `container.go` → infraestructura compartida + orquestación de módulos
- `<dominio>_module.go` → cableado completo de un dominio (repo → service → handler → routes)

## Ejemplo de uso

```go
// main.go
func main() {
    c := bootstrap.NewContainer()

    r := gin.Default()
    router.SetupRoutes(r, c.Registrars)
    r.Run(":8080")
}
```

## Extendiendo el container

Cuando agregues un nuevo dominio, por ejemplo `Invoice`:

**1. Crear `invoice_module.go`** — encapsula todo el cableado del dominio:

```go
// bootstrap/invoice_module.go
func newInvoiceModule(db *sql.DB) router.RouteRegistrar {
    invoiceRepo    := mysql.NewInvoiceRepository(db)
    invoiceService := service.NewInvoiceService(invoiceRepo)
    invoiceHandler := handler.NewInvoiceHandler(invoiceService)
    return routes.NewInvoiceRoutes(invoiceHandler)
}
```

**2. Agregar una línea en `container.go`**:

```go
Registrars: []router.RouteRegistrar{
    newInvoiceModule(db),           // ← única línea nueva
},
```

`container.go` crece exactamente una línea por módulo nuevo.
