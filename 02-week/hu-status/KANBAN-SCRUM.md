# Buenas prácticas y principios para el uso de Kanban y Scrum

## 1. Introducción

La gestión ágil de proyectos de software requiere mecanismos que permitan organizar el trabajo, controlar el flujo de actividades, responder al cambio y mantener una entrega continua de valor. Dentro de este contexto, **Scrum** y **Kanban** son dos de los enfoques más utilizados.

Aunque ambos se relacionan con Agile, no son equivalentes ni resuelven exactamente el mismo problema.

**Scrum** proporciona un marco estructurado para desarrollar productos complejos mediante iteraciones denominadas **Sprints**, estableciendo responsabilidades, eventos y artefactos.

**Kanban**, por otro lado, es un método orientado a la **gestión y optimización del flujo de trabajo**, utilizando visualización, límites de trabajo en progreso y métricas de flujo.

La utilización correcta de cualquiera de los dos requiere comprender que las herramientas, tableros y reuniones no constituyen por sí mismas agilidad. Un equipo puede tener un tablero perfectamente configurado y seguir trabajando de manera completamente ineficiente. La metodología debe utilizarse para mejorar la capacidad del equipo de **entregar valor, detectar problemas y adaptarse**.

Este documento presenta los principios, prácticas, responsabilidades, métricas, errores frecuentes y criterios de aplicación de Kanban y Scrum, además de analizar cómo pueden combinarse en un mismo proyecto.

---

# 2. Fundamentos de Agile

Scrum y Kanban se utilizan frecuentemente dentro de entornos ágiles, por lo que es importante comprender primero los principios generales sobre los cuales se construyen estas prácticas.

## 2.1. Entrega frecuente de valor

El objetivo principal de un proyecto de software no es producir tareas completadas, documentación o código por sí mismos.

El objetivo es producir **incrementos funcionales que generen valor**.

Por lo tanto:

```text
Actividad realizada ≠ Valor entregado
```

Una tarea puede estar técnicamente terminada y no aportar ningún beneficio relevante al producto.

Un enfoque correcto prioriza:

```text
Necesidad
   ↓
Trabajo
   ↓
Resultado
   ↓
Valor
```

---

## 2.2. Transparencia

La información necesaria para tomar decisiones debe ser visible para las personas involucradas.

La transparencia permite conocer:

* Qué trabajo existe.
* Qué trabajo tiene mayor prioridad.
* Qué está actualmente en progreso.
* Qué está bloqueado.
* Qué está terminado.
* Qué riesgos existen.
* Qué problemas están afectando el flujo.
* Qué valor se ha entregado.

Un sistema que oculta problemas no los elimina. Simplemente consigue que aparezcan más tarde y con intereses.

---

## 2.3. Inspección

El equipo debe revisar periódicamente:

* El producto.
* El progreso.
* La calidad.
* El flujo.
* Los resultados.
* El proceso de trabajo.

La inspección permite detectar desviaciones antes de que se conviertan en problemas mayores.

---

## 2.4. Adaptación

Cuando una inspección demuestra que algo no está funcionando, el proceso debe modificarse.

```text
Ejecutar
   ↓
Medir
   ↓
Inspeccionar
   ↓
Detectar desviación
   ↓
Adaptar
   ↓
Ejecutar nuevamente
```

La adaptación constituye una de las características fundamentales de los enfoques ágiles.

---

## 2.5. Mejora continua

La mejora continua implica que el proceso de trabajo debe evolucionar constantemente.

No se trata de cambiar todo cada semana.

Se trata de:

1. Identificar un problema.
2. Comprender su causa.
3. Proponer una mejora.
4. Aplicarla.
5. Medir su resultado.
6. Mantenerla, modificarla o descartarla.

---

# 3. KANBAN

## 3.1. Definición

**Kanban** es un método de gestión del trabajo basado en la visualización y optimización del flujo.

Su propósito principal es permitir que el trabajo avance de manera **predecible, eficiente y sostenible** desde la solicitud hasta la entrega.

Kanban no exige necesariamente Sprints, roles específicos ni iteraciones temporales.

Su unidad fundamental es el **elemento de trabajo** y su recorrido a través de un sistema.

Ejemplo:

```text
┌─────────┐
│ BACKLOG │
└────┬────┘
     ↓
┌─────────┐
│  TODO   │
└────┬────┘
     ↓
┌─────────────┐
│ IN PROGRESS │
└──────┬──────┘
       ↓
┌─────────┐
│ REVIEW  │
└────┬────┘
     ↓
┌─────────┐
│  DONE   │
└─────────┘
```

---

# 4. Principios fundamentales de Kanban

Kanban se basa en varios principios que deben mantenerse incluso cuando el tablero o las herramientas utilizadas cambien.

## 4.1. Visualizar el trabajo

El primer principio consiste en representar visualmente el trabajo.

Un tablero Kanban debe mostrar el estado real de los elementos.

Ejemplo:

```text
BACKLOG       TODO          IN PROGRESS       REVIEW        DONE

Task 8        Task 5        Task 2             Task 1       Task 0
Task 9        Task 6        Task 3
              Task 7
```

La visualización permite identificar rápidamente:

* Acumulaciones.
* Cuellos de botella.
* Tareas bloqueadas.
* Exceso de WIP.
* Falta de trabajo.
* Problemas de capacidad.

---

## 4.2. Limitar el Work in Progress

**WIP (Work in Progress)** representa el trabajo que actualmente se encuentra en proceso.

Uno de los principios más importantes de Kanban consiste en **limitar el WIP**.

Ejemplo:

```text
IN PROGRESS
WIP LIMIT = 3

┌─────────┐
│ Task A  │
├─────────┤
│ Task B  │
├─────────┤
│ Task C  │
└─────────┘
```

Mientras existan tres tareas en progreso, el equipo debería priorizar su finalización antes de iniciar nuevas tareas.

### ¿Por qué?

Porque iniciar demasiadas tareas simultáneamente genera:

* Cambio constante de contexto.
* Mayor tiempo de finalización.
* Mayor cantidad de tareas incompletas.
* Más bloqueos.
* Mayor dificultad de seguimiento.

La optimización del trabajo no consiste en tener a todas las personas ocupadas todo el tiempo.

Consiste en **hacer que el trabajo termine**.

---

# 5. Ley de Little

Una relación importante para comprender el comportamiento de un sistema Kanban es la **Ley de Little**:

```text
WIP = Throughput × Cycle Time
```

Por lo tanto:

```text
Cycle Time = WIP / Throughput
```

Donde:

* **WIP** = trabajo en progreso.
* **Throughput** = cantidad de elementos completados por unidad de tiempo.
* **Cycle Time** = tiempo necesario para completar un elemento.

Ejemplo:

```text
WIP = 10 tareas
Throughput = 5 tareas/semana

Cycle Time = 10 / 5
           = 2 semanas
```

Esto permite comprender por qué aumentar indefinidamente el número de tareas en progreso no necesariamente acelera la entrega.

---

# 6. Gestionar el flujo

Kanban busca optimizar el flujo completo y no únicamente una etapa individual.

Un problema común ocurre cuando una columna recibe trabajo más rápido de lo que puede procesarlo.

```text
DEVELOPMENT       REVIEW

████████████      ██
████████████      ██
████████████      ██
```

En este caso, desarrollo está produciendo trabajo más rápido de lo que revisión puede procesarlo.

El resultado será:

```text
Acumulación
   ↓
Mayor WIP
   ↓
Mayor Cycle Time
   ↓
Retrasos
```

La solución no necesariamente consiste en aumentar la velocidad de desarrollo.

Puede ser necesario aumentar la capacidad de revisión.

---

# 7. Políticas explícitas

Cada transición entre estados debería tener criterios claros.

Ejemplo:

```text
TODO
↓
Requisitos definidos
↓
IN PROGRESS
↓
Desarrollo iniciado
↓
REVIEW
↓
Código implementado + pruebas ejecutadas
↓
DONE
↓
Criterios de aceptación cumplidos
```

Estas reglas constituyen políticas explícitas del sistema.

---

# 8. Clases de servicio

Kanban puede utilizar diferentes clases de servicio para representar distintas necesidades.

### Standard

Trabajo normal.

### Expedite

Trabajo excepcionalmente urgente.

Debe utilizarse con mucha moderación.

### Fixed Date

Trabajo asociado a una fecha límite específica.

### Intangible

Trabajo cuyo beneficio puede ser técnico o preventivo.

Ejemplos:

* Refactorización.
* Actualización de dependencias.
* Mejoras de infraestructura.
* Reducción de deuda técnica.

Una mala práctica es clasificar todo como urgente.

Si todo es urgente, nada tiene realmente prioridad.

---

# 9. Métricas de Kanban

Kanban utiliza principalmente métricas relacionadas con el flujo.

## 9.1. Lead Time

Tiempo desde que una necesidad entra al sistema hasta que se entrega.

```text
Solicitud ───────────────────────► Entrega
             Lead Time
```

---

## 9.2. Cycle Time

Tiempo desde que comienza activamente el trabajo hasta que termina.

```text
Inicio ─────────────► DONE
       Cycle Time
```

---

## 9.3. Throughput

Cantidad de elementos terminados durante un periodo.

Ejemplo:

```text
Semana 1 → 8 tareas
Semana 2 → 11 tareas
Semana 3 → 9 tareas
```

---

## 9.4. WIP

Cantidad de elementos actualmente en proceso.

---

## 9.5. Cumulative Flow Diagram

El **Cumulative Flow Diagram (CFD)** permite observar la evolución del trabajo en las diferentes etapas.

Sirve para detectar:

* Cuellos de botella.
* Acumulaciones.
* Cambios en el throughput.
* Problemas de flujo.

---

# 10. Buenas prácticas de Kanban

Un sistema Kanban debería:

1. Visualizar todo el trabajo relevante.
2. Mantener límites WIP.
3. Priorizar explícitamente.
4. Hacer visibles los bloqueos.
5. Definir políticas claras.
6. Medir Cycle Time.
7. Medir Throughput.
8. Revisar periódicamente el flujo.
9. Atacar cuellos de botella.
10. Evitar multitarea excesiva.
11. Mantener las tareas suficientemente pequeñas.
12. Mejorar progresivamente el sistema.

---

# 11. SCRUM

## 11.1. Definición

**Scrum** es un marco de trabajo para desarrollar, entregar y mantener productos complejos.

El trabajo se organiza en ciclos de duración fija denominados **Sprints**.

Cada Sprint debe producir un **Incremento** que aporte valor al producto.

```text
PRODUCT BACKLOG
       ↓
SPRINT PLANNING
       ↓
SPRINT
       ↓
INCREMENT
       ↓
SPRINT REVIEW
       ↓
RETROSPECTIVE
       ↓
NEXT SPRINT
```

---

# 12. Componentes de Scrum

Scrum define tres responsabilidades principales, tres artefactos y cinco eventos formales.

## 12.1. Responsabilidades

### Product Owner

Responsable de maximizar el valor del producto y gestionar eficazmente el **Product Backlog**.

Entre sus responsabilidades se encuentran:

* Definir y comunicar el Product Goal.
* Ordenar el Product Backlog.
* Asegurar que los elementos sean comprensibles.
* Representar las necesidades de los stakeholders.
* Tomar decisiones relacionadas con el valor del producto.

---

### Scrum Master

Responsable de ayudar a establecer Scrum y mejorar la efectividad del equipo.

Sus funciones incluyen:

* Facilitar Scrum.
* Ayudar a eliminar impedimentos.
* Promover la mejora continua.
* Facilitar eventos cuando sea necesario.
* Ayudar a la organización a comprender Scrum.

El Scrum Master no es el jefe del equipo.

---

### Developers

Son las personas responsables de crear el incremento.

Sus responsabilidades incluyen:

* Crear el Sprint Backlog.
* Adaptar diariamente el plan.
* Mantener la calidad.
* Cumplir la Definition of Done.
* Crear incrementos utilizables.

---

# 13. Artefactos de Scrum

## 13.1. Product Backlog

Lista ordenada y emergente de todo aquello que puede ser necesario para mejorar el producto.

```text
PRODUCT BACKLOG

1. Login
2. Gestión de usuarios
3. Sistema de pagos
4. Reportes
5. Notificaciones
```

No es simplemente una lista de deseos.

Debe estar ordenado según valor, necesidad, riesgo y contexto.

---

## 13.2. Sprint Backlog

Conjunto de elementos seleccionados para el Sprint junto con el plan necesario para alcanzar el Sprint Goal.

```text
SPRINT GOAL
     ↓
"Permitir registrar pedidos"

SPRINT BACKLOG
├── Crear endpoint
├── Crear modelo
├── Implementar validación
├── Crear interfaz
└── Crear pruebas
```

---

## 13.3. Increment

Resultado concreto producido durante el Sprint que cumple la **Definition of Done**.

Un incremento debe encontrarse en condiciones de ser utilizado.

---

# 14. Sprint Goal

El **Sprint Goal** establece el propósito principal del Sprint.

Debe representar un resultado y no simplemente una colección de tareas.

### Incorrecto

```text
Completar 15 tickets.
```

### Correcto

```text
Permitir que un cliente pueda crear y consultar
el estado de un pedido.
```

El Sprint Goal proporciona dirección cuando aparecen cambios o dificultades durante el Sprint.

---

# 15. Eventos de Scrum

## 15.1. Sprint

Es el contenedor de todos los demás eventos.

Durante el Sprint se realiza el trabajo necesario para producir el incremento.

---

## 15.2. Sprint Planning

Se determina:

1. ¿Por qué este Sprint es valioso?
2. ¿Qué puede realizarse?
3. ¿Cómo se realizará el trabajo?

El resultado incluye:

```text
Sprint Goal
+
Sprint Backlog
```

---

## 15.3. Daily Scrum

Evento breve utilizado por los Developers para inspeccionar el progreso hacia el Sprint Goal y adaptar el plan.

No debería convertirse en un reporte individual dirigido al jefe.

El foco debe ser:

```text
¿Dónde estamos?
¿Qué está bloqueando el objetivo?
¿Qué debemos adaptar?
```

---

## 15.4. Sprint Review

Se inspecciona el resultado del Sprint junto con los stakeholders relevantes.

Se analiza:

* Qué se construyó.
* Qué cambió.
* Qué feedback existe.
* Qué debería hacerse posteriormente.

No debe limitarse a una demostración superficial.

---

## 15.5. Sprint Retrospective

El equipo inspecciona su forma de trabajar.

Preguntas útiles:

```text
¿Qué funcionó?
¿Qué no funcionó?
¿Por qué?
¿Qué podemos cambiar?
¿Qué acción concreta tomaremos?
```

Una retrospectiva sin acciones posteriores es básicamente terapia grupal con post-its.

---

# 16. Definition of Done

La **Definition of Done (DoD)** es una descripción formal del estado requerido para considerar un incremento terminado.

Ejemplo:

```text
Definition of Done

- Código implementado.
- Pruebas unitarias realizadas.
- Pruebas de integración realizadas cuando corresponda.
- Code review completado.
- Criterios de aceptación cumplidos.
- Documentación actualizada cuando sea necesaria.
- Sin errores críticos conocidos.
- Integración realizada.
- Cumple estándares técnicos del proyecto.
```

La DoD establece un estándar de calidad común.

Una tarea que todavía requiere pruebas, revisión o correcciones no debería considerarse terminada.

---

# 17. Buenas prácticas de Scrum

## 17.1. Mantener el Product Backlog ordenado

Los elementos de mayor prioridad deben encontrarse suficientemente preparados.

---

## 17.2. Crear elementos pequeños

Una historia demasiado grande genera:

* Incertidumbre.
* Mayor dificultad de estimación.
* Mayor riesgo.
* Menor visibilidad.

Es preferible dividirla verticalmente.

```text
Historia grande
      ↓
┌──────────────┐
│ Crear pedido │
└──────────────┘
      ↓
┌──────────────┬──────────────┬──────────────┐
│ Crear pedido │ Validar      │ Consultar    │
│              │ pedido       │ estado       │
└──────────────┴──────────────┴──────────────┘
```

---

## 17.3. Definir criterios de aceptación

Cada Product Backlog Item debe tener condiciones que permitan determinar objetivamente si cumple la necesidad.

Ejemplo:

```text
Dado que el usuario está autenticado

Cuando crea un pedido válido

Entonces el sistema debe registrar el pedido
y asignarle un identificador único.
```

---

## 17.4. Evitar sobrecargar el Sprint

El Sprint Backlog debe reflejar una cantidad de trabajo razonable.

Agregar constantemente trabajo sin considerar la capacidad del equipo destruye la previsibilidad.

---

## 17.5. Mantener la calidad durante el Sprint

La calidad no debería acumularse para el final.

```text
Implementar
   ↓
Probar
   ↓
Revisar
   ↓
Integrar
   ↓
Completar
```

No:

```text
Implementar 40 cosas
       ↓
Probar todo al final
       ↓
Descubrir 23 problemas
       ↓
La humanidad llora
```

---

# 18. Estimación

La estimación en Scrum busca ayudar a comprender la complejidad relativa del trabajo.

Puede utilizarse **Story Points**.

Los Story Points normalmente consideran:

* Complejidad.
* Incertidumbre.
* Esfuerzo relativo.
* Riesgo.

No representan directamente horas.

Ejemplo:

```text
Historia A → 1 punto
Historia B → 3 puntos
Historia C → 5 puntos
Historia D → 8 puntos
```

Una historia de 8 puntos representa un trabajo considerablemente más complejo que una de 3, pero no significa automáticamente "8 horas".

---

# 19. Velocity

La **Velocity** representa la cantidad de Story Points completados durante un Sprint.

Ejemplo:

```text
Sprint 1 → 21
Sprint 2 → 24
Sprint 3 → 23
Sprint 4 → 26
```

Promedio aproximado:

```text
Velocity ≈ 23.5
```

La velocity puede utilizarse para planificación, pero no debería convertirse en un objetivo de rendimiento individual.

### Mala práctica

```text
"Tenemos que aumentar la velocity un 20%."
```

Esto incentiva inflar estimaciones.

La velocity es una **métrica de planificación**, no una métrica de productividad personal.

---

# 20. Kanban vs Scrum

| Característica       | Kanban           | Scrum                       |
| -------------------- | ---------------- | --------------------------- |
| Naturaleza           | Método de flujo  | Framework                   |
| Flujo                | Continuo         | Iterativo                   |
| Sprint               | No requerido     | Obligatorio                 |
| Roles específicos    | No               | Sí                          |
| Eventos formales     | No obligatorios  | Sí                          |
| WIP                  | Límite explícito | Capacidad del Sprint        |
| Cambios              | Continuos        | Se protege el Sprint Goal   |
| Priorización         | Continua         | Product Backlog             |
| Métricas principales | Flow metrics     | Métricas de Sprint/producto |
| Cycle Time           | Fundamental      | Útil                        |
| Throughput           | Fundamental      | Útil                        |
| Velocity             | No necesaria     | Puede utilizarse            |
| Mejor escenario      | Flujo continuo   | Producto complejo           |
| Adaptación           | Continua         | Por Sprint                  |

---

# 21. Diferencia conceptual fundamental

La diferencia puede resumirse de esta manera:

```text
KANBAN
──────

¿Cómo hacemos que el trabajo fluya mejor?

Visualizar
    ↓
Limitar WIP
    ↓
Eliminar bloqueos
    ↓
Optimizar flujo
    ↓
Reducir Cycle Time


SCRUM
─────

¿Cómo desarrollamos un producto complejo
mediante ciclos de inspección y adaptación?

Product Goal
    ↓
Sprint Goal
    ↓
Sprint
    ↓
Increment
    ↓
Review
    ↓
Retrospective
    ↓
Adaptación
```

---

# 22. Scrumban

Kanban y Scrum no tienen que utilizarse como alternativas completamente aisladas.

Un equipo puede mantener la estructura de Scrum y adoptar prácticas de Kanban.

Esto suele denominarse **Scrumban**.

Ejemplo:

```text
SCRUM
├── Product Backlog
├── Sprint
├── Sprint Goal
├── Review
└── Retrospective

+

KANBAN
├── Visualización
├── WIP limits
├── Gestión de bloqueos
├── Cycle Time
└── Flow metrics
```

La combinación puede resultar útil cuando el equipo necesita la estructura de Scrum pero también necesita mejorar significativamente su flujo.

---

# 23. Gestión de bloqueos

Un bloqueo debe hacerse visible inmediatamente.

Ejemplo:

```text
IN PROGRESS

[Implementar pagos]
       ↓
    BLOCKED
       ↓
Dependencia externa
```

Un bloqueo prolongado puede provocar:

```text
Bloqueo
   ↓
WIP aumenta
   ↓
Cycle Time aumenta
   ↓
Throughput disminuye
   ↓
Entrega se retrasa
```

Buenas prácticas:

* Identificar el bloqueo.
* Registrar su causa.
* Asignar responsabilidad para resolverlo.
* Evitar que el equipo siga acumulando trabajo.
* Revisar bloqueos recurrentes.
* Eliminar la causa raíz cuando sea posible.

---

# 24. Gestión de dependencias

En proyectos de software, algunas tareas dependen de otros equipos, servicios o componentes.

Ejemplo:

```text
Frontend
   │
   └────► API Backend
               │
               └────► Base de datos
                         │
                         └────► Servicio externo
```

Las dependencias deberían identificarse antes de que bloqueen el trabajo.

Buenas prácticas:

* Detectar dependencias durante la planificación.
* Documentar responsables.
* Reducir acoplamiento.
* Crear contratos entre servicios.
* Utilizar mocks o stubs cuando sea apropiado.
* Automatizar pruebas de integración.

---

# 25. Gestión de riesgos

La agilidad no elimina los riesgos.

Permite detectarlos y reaccionar antes.

Los principales riesgos pueden ser:

* Técnicos.
* Funcionales.
* De integración.
* De infraestructura.
* De seguridad.
* De disponibilidad.
* De dependencia externa.
* De conocimiento.
* De calendario.

Una estrategia adecuada consiste en atacar primero las incertidumbres importantes.

```text
Alta incertidumbre
        ↓
Experimentar / validar
        ↓
Obtener información
        ↓
Reducir riesgo
```

---

# 26. Buenas prácticas para tableros digitales

Independientemente de utilizar Jira, GitHub Projects, Azure DevOps, Trello u otra herramienta, el tablero debería:

### Tener estados claros

```text
BACKLOG
TODO
IN PROGRESS
REVIEW
DONE
```

### Evitar estados redundantes

No es necesario crear:

```text
Started
Working
Development
Coding
Implementation
In Progress
```

Son seis maneras de decir "alguien está trabajando en esto".

### Mostrar información relevante

Cada tarjeta debería permitir comprender rápidamente:

* Qué debe hacerse.
* Quién trabaja en ella.
* Prioridad.
* Estado.
* Dependencias.
* Criterios de aceptación.
* Bloqueos.

---

# 27. Integración con Git

La metodología debe conectarse con el proceso técnico de desarrollo.

Un flujo recomendado:

```text
Backlog Item
     ↓
Task / Story
     ↓
Branch
     ↓
Commit
     ↓
Pull Request
     ↓
Code Review
     ↓
CI / Tests
     ↓
Merge
     ↓
Deploy
     ↓
Done
```

Esto permite conectar el trabajo de gestión con el trabajo técnico.

---

# 28. Integración continua

Las prácticas ágiles funcionan mejor cuando el equipo puede integrar cambios frecuentemente.

Un pipeline puede ser:

```text
Push
 ↓
Build
 ↓
Lint
 ↓
Unit Tests
 ↓
Integration Tests
 ↓
Security Checks
 ↓
Deploy
```

La automatización reduce errores manuales y proporciona feedback rápido.

---

# 29. Definition of Ready

Aunque no constituye un requisito formal de Scrum, algunos equipos utilizan una **Definition of Ready (DoR)** para determinar cuándo un elemento está suficientemente preparado para ser trabajado.

Ejemplo:

```text
Definition of Ready

- Objetivo claro.
- Descripción comprensible.
- Criterios de aceptación definidos.
- Dependencias identificadas.
- Tamaño razonable.
- Prioridad establecida.
- Información suficiente para comenzar.
```

Debe utilizarse como una ayuda, no como una barrera burocrática.

---

# 30. Errores frecuentes en Scrum

## 30.1. Convertir Scrum en cascada de dos semanas

Un equipo puede terminar haciendo:

```text
Semana 1 → análisis
Semana 2 → desarrollo
Semana 3 → pruebas
Semana 4 → integración
```

y simplemente llamarlo "Sprints".

Eso no lo convierte mágicamente en Scrum.

---

## 30.2. Medir productividad por cantidad de tareas

Cerrar muchos tickets pequeños no necesariamente significa entregar más valor.

---

## 30.3. Utilizar velocity como KPI individual

Esto genera comportamientos perversos:

```text
Estimación inflada
      ↓
Mayor velocity
      ↓
Métrica "mejor"
      ↓
Producto no necesariamente mejor
```

---

## 30.4. Daily Scrum como reporte al jefe

El Daily Scrum debe ayudar a los Developers a inspeccionar y adaptar su plan.

---

## 30.5. Ignorar la retrospectiva

Si los mismos problemas aparecen Sprint tras Sprint, el equipo no está mejorando.

---

# 31. Errores frecuentes en Kanban

## 31.1. Utilizar el tablero como simple lista de tareas

Un tablero Kanban debe representar el **flujo**, no solamente almacenar tareas.

---

## 31.2. No establecer límites WIP

Sin límites WIP se pierde uno de los principales mecanismos de Kanban.

---

## 31.3. Tener demasiadas columnas

La complejidad visual debe reflejar el proceso real.

---

## 31.4. Ignorar métricas

Sin medición resulta difícil saber si el flujo realmente está mejorando.

---

## 31.5. Permitir que todo sea urgente

Las excepciones permanentes destruyen las políticas del sistema.

---

# 32. Principios de diseño de un buen sistema de trabajo

Un sistema basado en Scrum, Kanban o ambos debería buscar:

```text
                    VALOR
                      │
        ┌─────────────┼─────────────┐
        ↓             ↓             ↓
   TRANSPARENCIA   CALIDAD       FLUJO
        │             │             │
        └─────────────┼─────────────┘
                      ↓
                 INSPECCIÓN
                      ↓
                 ADAPTACIÓN
                      ↓
              MEJORA CONTINUA
```

Estos principios permiten que las herramientas sean secundarias frente al objetivo real.

---

# 33. Scrum y Kanban en un proyecto de software

Para un proyecto académico o profesional de desarrollo de software puede utilizarse una combinación práctica:

```text
PRODUCT BACKLOG
       ↓
Priorización
       ↓
SPRINT PLANNING
       ↓
SPRINT BACKLOG
       ↓
┌─────────────────────────────────┐
│           KANBAN BOARD          │
│                                 │
│ TODO → DEV → REVIEW → TEST → DONE
│          │                      │
│          └── WIP LIMIT          │
└─────────────────────────────────┘
       ↓
INCREMENT
       ↓
SPRINT REVIEW
       ↓
RETROSPECTIVE
       ↓
MEJORAS
```

En este modelo:

* **Scrum** proporciona la estructura temporal.
* **Kanban** proporciona mecanismos para optimizar el flujo.
* **Git** controla la evolución del código.
* **CI/CD** automatiza validación y entrega.
* **Métricas** permiten evaluar el sistema.

---

# 34. Métricas recomendadas

No es necesario medir absolutamente todo.

Las métricas deben utilizarse para responder preguntas concretas.

| Pregunta                                 | Métrica      |
| ---------------------------------------- | ------------ |
| ¿Cuánto tardamos en entregar?            | Lead Time    |
| ¿Cuánto tarda el desarrollo?             | Cycle Time   |
| ¿Cuánto terminamos?                      | Throughput   |
| ¿Cuánto trabajo tenemos abierto?         | WIP          |
| ¿Cuánto completamos por Sprint?          | Velocity     |
| ¿Cuánto trabajo falla?                   | Defect Rate  |
| ¿Cuánto tiempo estamos bloqueados?       | Blocked Time |
| ¿Qué tan rápido recuperamos el servicio? | MTTR         |

Las métricas deben utilizarse para **mejorar el sistema**, no para castigar personas.

---

# 35. Principio de optimización

Una de las ideas más importantes en Agile es que **optimizar una parte del sistema no necesariamente optimiza el sistema completo**.

Ejemplo:

```text
Backend
  ↓
Produce 20 tareas/día

QA
  ↓
Procesa 8 tareas/día
```

Aumentar todavía más la productividad del Backend puede empeorar el sistema.

Resultado:

```text
20 tareas producidas
        ↓
8 revisadas
        ↓
12 acumuladas
        ↓
Mayor WIP
        ↓
Mayor Lead Time
```

La optimización debe centrarse en el **cuello de botella**.

---

# 36. Principios para equipos de desarrollo

Un equipo que utilice Scrum, Kanban o una combinación debería adoptar las siguientes prácticas:

### 1. Trabajar sobre prioridades reales

El trabajo debe estar relacionado con los objetivos del producto.

### 2. Mantener tareas pequeñas

Las unidades pequeñas proporcionan feedback más rápido.

### 3. Evitar multitarea

Terminar antes de comenzar más trabajo.

### 4. Automatizar

Automatizar:

* Pruebas.
* Builds.
* Deployments.
* Validaciones.
* Linting.
* Análisis estático.

### 5. Integrar frecuentemente

Evitar ramas extremadamente largas y grandes cambios difíciles de integrar.

### 6. Revisar código

Utilizar code reviews para mejorar calidad y compartir conocimiento.

### 7. Medir resultados

Utilizar métricas para comprender el sistema.

### 8. Atacar causas raíz

No limitarse a solucionar repetidamente los mismos síntomas.

---

# 37. Principios que NO deben confundirse

## Agile ≠ Scrum

Agile es una filosofía y conjunto de principios.

Scrum es un framework.

---

## Scrum ≠ Jira

Jira es una herramienta.

Puede utilizarse para implementar Scrum, Kanban u otros procesos.

---

## Kanban ≠ tablero

El tablero es una herramienta de visualización.

Kanban implica además gestión del flujo, límites WIP, políticas explícitas y mejora continua.

---

## Story Points ≠ horas

Son una medida relativa de complejidad/esfuerzo.

---

## Velocity ≠ productividad

La velocity sirve principalmente para planificación y previsibilidad.

---

## Done ≠ "ya escribí el código"

Done debe cumplir la Definition of Done.

---

# 38. Reglas prácticas para el equipo

Como guía operativa, un equipo puede establecer las siguientes reglas:

```text
1. Todo trabajo debe estar visible.
2. Toda tarea debe tener una prioridad.
3. No comenzar trabajo innecesario.
4. Respetar los límites WIP.
5. Priorizar terminar antes que comenzar.
6. Los bloqueos deben hacerse visibles.
7. Las historias deben tener criterios de aceptación.
8. El código debe cumplir la Definition of Done.
9. Los cambios deben integrarse frecuentemente.
10. Los problemas recurrentes deben generar acciones de mejora.
11. Las métricas sirven para mejorar el proceso, no para castigar personas.
12. El proceso debe adaptarse cuando la evidencia demuestra que no funciona.
```

---

# 39. Resumen conceptual

```text
                         AGILE
                           │
          ┌────────────────┴────────────────┐
          │                                 │
        SCRUM                              KANBAN
          │                                 │
    Iteraciones                         Flujo continuo
    Sprints                             Visualización
    Sprint Goal                         WIP limits
    Increment                           Flow metrics
    Review                              Cycle Time
    Retrospective                       Throughput
          │                                 │
          └────────────────┬────────────────┘
                           ↓
                   ENTREGA DE VALOR
                           ↓
                    INSPECCIÓN
                           ↓
                     ADAPTACIÓN
                           ↓
                  MEJORA CONTINUA
```

---

# 40. Conclusiones

Scrum y Kanban son herramientas diferentes para abordar la gestión del trabajo y el desarrollo de productos.

**Scrum** resulta especialmente útil cuando se necesita una estructura iterativa clara, objetivos por Sprint, responsabilidades definidas y ciclos regulares de inspección y adaptación.

**Kanban** resulta especialmente útil cuando se necesita gestionar un flujo continuo, visualizar el trabajo, limitar el WIP y optimizar el tiempo de entrega.

Ambos enfoques pueden complementarse.

La elección correcta no debería depender de cuál metodología parece más moderna o de cuál tablero se ve más bonito. Debe depender de **las características del producto, el flujo de trabajo, la incertidumbre, las necesidades del equipo y el tipo de entrega que se necesita realizar**.

El principio fundamental puede resumirse en:

```text
VISUALIZAR
     ↓
PRIORIZAR
     ↓
LIMITAR EL TRABAJO
     ↓
ENTREGAR VALOR
     ↓
MEDIR
     ↓
INSPECCIONAR
     ↓
ADAPTAR
     ↓
MEJORAR
     ↺
```

El propósito de Scrum y Kanban no es hacer que el equipo parezca ágil.

El propósito es construir un sistema de trabajo capaz de **entregar valor de manera sostenible, mantener la calidad, responder al cambio y mejorar continuamente**.
