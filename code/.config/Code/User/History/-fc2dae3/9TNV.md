# Java

## Cos'è Java

- Linguaggio del 1995 orientato agli oggetti
- Cross-platform
- Sicuro - non permette manomissioni
- Robusto - molti errori rimossi a tempo di compilazione
- Multi-threaded - più attività in contemporanea
- Interpretato - byte code tradotto "al volo" in istruzioni macchina native
- Alte prestazioni - compilatore JIT (Just-In-Time)
- Enterprise, Web e Mobile
- WORA (Write Once, Run Anywhere)

## Tipi di dati di base o primitivi

I tipi di dati di base sono built-in, ovvero predefiniti nel linguaggio.

Un tipo di dati è un insieme di valori e di operazioni definiti su tali valori.

- Interi
- Reali
- Booleani
- Caratteri
- Stringhe

## Dichiarazione vs assegnazione vs inizializzazione di variabili

```java
// dichiarazione
int a;

// dichiarazione e inizializzazione
int b = 5;

// assegnazione
a = 10;
```

## Conversioni di tipi

### Conversione esplicita

Utilizzando un metodo che prende in ingresso un argomento di un tipo e restituisce un altro tipo:

```java
String text = "2";
Int number = Integer.ParseInt(text);
```

### Cast esplicito

Anteponendo il tipo desiderato tra parentesi:

```java
int a = (int) 2.71828;
```

### Cast implicito

Se il tipo di partenza è **meno preciso, Java converte automaticamente il valore** al tipo più preciso:

```java
double d = 2; // OK - cast implicito (eseguito automaticamente - no perdita di dati)

int d = 2.5; // KO - ERROR (perdita di dati/precisione)

int d = (int) 2.5; // Fix - cast esplicito (azione esplicita/consapevole della perdita)
```

## Classi e oggetti

- Una classe è un pezzo di codice sorgente di un programma che descrive un particolare tipo di oggetto
- Le classi vengono definite dal programmatore
- La classe definisce un prototipo astratto per gli oggetti di un particolare tipo
- Ne definisce la struttura in termini di **Attributi** (stato degli ogetti) e **Metodi** (comportamento degli ogetti)
- Un oggetto è un'istanza (un esemplare) di una classe
- Un programma può creare e usare uno o più oggetti (istanze) della classe

### Campo

- Un campo (detto anche **variabile di istanza**) costituisce la **memoria privata** di un oggetto
- Ogni campo ha un tipo di dati (es. contatore intero)
- Ogni campo ha un nome fornito dal programmatore

```java
// Dichiarazione di un campo

// private: accesso privato
// static: condiviso da tutti gli oggetti della classe
// final: campo costante

// private [static] [final] tipo_di_dati nome;
private static final double PI = 3.14;
```

### Metodo

- Un metodo è tipicamente pubblico, ovvero visibile a tutti
- Il nome di un metodo è scritto in camelCase

```java
// tipo_di_dati: tipo del valore restituito dal metodo
/*
public tipo_di_dati nomeDelMetodo(tipo_di_dati_1 param1, ...) {
    istruzioni
}
*/

public double getPI() {
    return 3.14;
}

// void: nessun valore restituito

public void helloWorld() {
    System.out.println("Hello World");
}
```

### Costruttori

- Metodi (funzioni) per la creazione degli oggetti di una classe
- Possiedono sempre lo stesso nome della classe
- Inizializzano i campi dell'oggetto
- Possono prendere zero o più parametri
- Non hanno valori di uscita (ma non specificano `void`)
- Una classe può avere più costruttori che differiscono dal numero e tipo di parametri
- Non è obbligatorio specificare un costruttore (Java ne crea uno di default senza parametri)
- Un oggetto della classe viene creato con l'operatore `new`

```java
public Counter() {
    value = 0;
}

public Counter(int initialValue) {
    value = initialValue;
}

Counter counter = new Counter();
```

### Variabili d'istanza (campi) vs variabili locali

- Le variabili di istanza (campi) sono variabili dell'oggetto (esistono per tutta la vita dell'oggetto)
- Le variabili locali sono definite all'interno di un metodo ed esistono fino al termine dell'esecuzione della chiamata al metodo

### Incapsulamento

- Il processo che nasconde i dettagli realizzativi, rendendo pubblica un'interfaccia
  - Dettagli realizzativi: campi e implementazioni
  - Interfaccia pubblica: metodi pubblici
- Si applica utilizzando le parole chiave **public** e **private**
- Utilizzato per nascondere le informazioni all'utente ("information hiding")
- Semplifica e modularizza il lavoro (funzionamento a scatola nera)
- Facilita il lavoro di gruppo e la maintenance
- Aiuta a rilevare errori

### Riferimenti ad oggetti

È importante tenere a mente la differenza della rappresentazione in memoria tra **valori primitivi** ed **oggetti**:

- Valori primitivi: memoria allocata automaticamente a tempo di compilazione
- Oggetti: memoria allocata durante l'esecuzione del programma (operatore `new`)

#### Riferimenti ad oggetti

- Per chi sa cosa sono i puntatori, i riferimenti sono ciò che ne è rimasto in Java
- Un riferimento è un indirizzo in memoria
- Assimilabili ai tipi di dati primitivi
- Quindi gli oggetti non sono mai memorizzati direttamente nelle variabili, ma solo mediante il loro riferimento

#### Stack e Heap

- Esistono due tipi di memoria: lo **heap** e lo **stack**
- Sullo stack vanno le variabili locali e i campi statici
- Sullo heap vanno le aree di memoria allocate per la creazione dinamica degli oggetti (oggetti e campi)

### Campi di classe: la parola chiave static

- I campi di una classe possono essere dichiarati `static`
- Un campo static è relativo all'intera classe, non al singolo oggetto istanziato
- Un campo static esiste in una sola locazione di memoria
- Viceversa, per ogni campo non static, esiste una locazione di memoria per ogni oggetto, allocata a seguito dell'istruzione `new`

### Metodi statici

- I metodi statici sono metodi di classe
- Non hanno accesso ai campi di istanza
- Ma hanno accesso ai campi di classe

### Classi wrapper

- Permettono di convertire i valori di un tipo primitivo in un oggetto
- Forniscono metodi di accesso e visualizzazione dei valori

#### Confrontare oggetti

- Confrontavamo i valori primitivi utilizzando operatori di confronto (`== != < <= > >=`)
- Ma ora, avendo degli oggetti, dobbiamo utilizzare metodi per il confronto:
  - `equals()`: restituisce `true` se e solo se l'oggetto in input è un intero di valore uguale al proprio
  - `compareTo()`: restituisce 0 se sono uguali, < 0 se il proprio valore è minore di quello in ingresso, > 0 altrimenti

#### Autoboxing e auto-unboxing

- L'autoboxing converte automaticamente un tipo primitivo al suo tipo wrapper associato

  ```java
  Integer k = 3;
  Integer[] array = {1,2,3,4,5};
  ```

- L'auto-unboxing converte automaticamente da un tipo wrapper all'equivalente tipo primitivo

  ```java
  Integer k = 3;
  Integer[] array = {1, 2, 3, 4, 5};

  int j = k;
  int n = array[j];
  ```

## Ereditarietà

- Concetto cardine della programmazione orientata agli oggetti
- Forma di riuso del software in cui una classe è creata:
  - "assorbendo" i membri di una classe esistente
  - aggiungendo nuove caratteristiche o migliorando quelle esistenti
- Aumenta la probabilità che il sistema sia implementato e manutenuto in maniera efficiente

### Che cosa si eredita?

- Una sottoclasse estende la superclasse
- La sottoclasse eredita i membri della superclasse
  - campi e metodi d'istanza secondo il livello di accesso specificato
- Inoltre la sottoclasse può:
  - aggiungere nuovi metodi e campi
  - ridefinire i metodi che eredita dall superclasse

```java
public class Forma {
    public void disegna();
}

public class Cerchio extends Forma {
  private double raggio;

  public Cerchio(double raggio) {
    this.raggio = raggio;
  }

  public double getRaggio() { return raggio; }
  public double getCirconferenza() { return 2 * Math.PI * raggio; }
}
```

### Classi astratte

- Una classe astratta (`abstract`) non può essere istanziata (non possono esistere oggetti per quella classe)
- Tipicamente verrà estesa da altre classi, che invece potranno essere istanziate

```java
public abstract class PersonaggioDisney {

}

public class Paperoga extends PersonaggioDisney {

}
```

### Metodi astratti

- Anche i metodi possono essere definiti astratti
  - Esclusivamente all'interno di una classe definita stratta
- **Non** forniscono l'implementazione per quel metodo
- Impongono alle sottoclassi non astratte di implementare il metodo

```java
public abstract class PersonaggioDisney {
    abstract void faPasticci();
}

public class Paperoga extends PersonaggioDisney {
    public void faPasticci() {
        System.out.println("$*!@#&^#!");
    }
}
```

### Parole chiave `this` e `super`

- Usate **obbligatoriamente** come nome di metodo nella prima riga del costruttore, permottono di:
  - `this`: richiamare un altro costruttore della stessa classe
  - `super`: richiamare un costruttore della superclasse
- Ogni sottoclasse deve esplicitamente definire un costruttore se la superclasse non fornisce un costruttore senza argomenti
- Quindi una sottoclasse, per poter essere istanziata, chiamerà **sempre** il costruttore della sua superclasse (anche se implicitamente)
  - In altre parole, prima che una sottoclasse possa essere istanziata e utilizzata, la sua superclasse deve essere istanziata e inizializzata correttamente

### Overriding vs overloading

- L'overriding consiste nel ridefinire (reimplementare) un metodo con la stessa intestazione ("segnatura") presente in una superlcasse
- L'overloading consiste nel creare un metodo con lo stesso nome, ma una intestazione diversa (diverso numero e/o tipo di paramentri)

#### Overriding

- Gli argomenti devono essere gli stessi
- I tipi di ritorno devono essere compatibili
- Non si può ridurre la visibilità (es. da public a private)

#### Overloading

- I tipi di ritorno possono essere diversi
  - non si può cambiare solo il tipo di ritorno
- Si può variare la visibilità in qualsiasi direzione

### In dettaglio sulla visibilità

- `Private`: visibile solo all'interno della classe
- `Public`: visibile a tutti
- `Default`: visibile all'interno di tutte le classi del package
- `Protected`: visibile all'interno di tutte le classi del package e delle sottoclassi (indipendentemente dal package)

### Is-a vs has-a

- Is-a rappresenta l'ereditarietà
  - Un oggetto di una sottoclasse può essere trattato come un oggetto della superclasse
  - Paperino è un PersonaggioDisney? - Sì
  - QuiQuoQua è un Paperino? - No
- Has-a rappresenta la composizione
  - Un oggetto contiene come membri riferimenti ad altri oggetti
  - Bagno contiene Vasca? - Sì
  - PersonaggioDisney contiene Paperino? - No

## Polimorfismo

- Insieme all'ereditarietà, altro concetto cardine della programmazione ad oggetti
- Polimorfismo = molte + forme
- Una variabile di un certo tipo A può contenere un riferimento ad una variabile del tipo A o di una sua qualiasi sottoclasse

  ```java
  Animale a = new Gatto();
  a = new Chihuahua();
  ```

- La selezione del metodo da chiamare avviene in base all'effettivo tipo dell'oggetto riferito dalla variabile
  - Questo si ottiene grazie al _binding dinamico_
  ```java
  Animale a = new Gatto();
  a.emettiVerso(); // Miaooo!
  a = new Chihuahua();
  a.emettiVerso(); // Bau Bau!
  ```

### Operatore `instanceof`

- L'operatore, applicato ad un oggetto e ad un nome di classe, restituisce true se l'oggetto è un tipo o sottotipo di quella classe

### Upcasting vs downcasting

#### Upcasting

- Posso sempre convertire senza cast esplicito un sottotipo ad un supertipo (upcasting)

```java
ImpiegatoStipendiato is1 = new ImpiegatoStipendiato("Mario", "imp1", 1500);
Impiegato i = is1;
```

- Con l'upcasting si "restringe" temporaneamente l'interfaccia dell'oggetto alla superclasse (perdita di dettaglio)

#### Downcasting

- A volte può essere necessario convertire un supertipo ad un sottotipo (downcasting)

```java
ImpiegatoStipendiato is1 = new ImpiegatoStipendiato("Mario", "imp1", 1500);
Impiegato i = is1;
ImpiegatoStipendiato is2 = (ImpiegatoStipendiato)i;
```

### La superclasse universale Object

- Tutte le classi in java ereditano direttamente o indirettamente dalla classe Object
- Quando si definisce una classe senza estenderne un'altra, questo è equivalente ad estendere Object
- La classe Object è una sorta di "minimo comun denominatore" di tutte le classi in Java

### Sovrascrivere metodi ereditati

- È possibile sovrascrivere metodi ereditati dalla superclasse utilizzando l'annotazione `@Override`
- `toString` è uno dei metodi che si ereditano dalla classe Object

### Metodi e classi final

- Con la parola chiave `abstract` **obblighiamo** i programmatori ad implementare certi metodi
- La parola chiave `final`, invece, **impedisce** ad altri programmatori di:
  - Creare sottoclassi - nonché di estendere la classe (se specificato di fronte a `class`)
  - Reimplementare (sovrascrivere) certi metodi (se specificato di fronte all'implementazione del metodo)

### La classe java.util.ArrayList

- Implementa una lista o un vettore ad accesso casuale
- Utilizza un array per implementare la lista ma:
  - La dimensione della lista può aumentare o diminuire
  - La classe fornisce metodi per svolgere le operazioni più importanti

```java
// Esempio di istanziazione

ArrayList<String> l = new ArrayList<String>();

/* La notazione <String> indica che utilizziamo una lista di stringhe e
   che la classe è generica, ovvero può essere utilizzata con tipi differenti
*/
```

## Interfacce e classi interne

### Le interfacce

- Le interfacce sono uno strumento per consentire a più classi di fornire ed implementare un insieme di metodi comuni
- Le interfacce definiscono e standardizzano l'interazione fra oggetti tramite un insieme limitato di operazioni
- Esse specificano solamente il **comportamento** che un certo oggetto deve presentare all'esterno, cioè cosa quell'oggetto può fare
- L'implementazione di tali operazioni, cioè come queste vengono tradotte e realizzate, rimane invece non definito
- **Le interfacce sono classi astratte al 100%** (se non definiamo metodi di default)
- Un'interfaccia è una classe che può contenere soltanto:
  - Costanti
  - Metodi astratti
  - Implementazione di default di metodi e metodi statici (Java >= 8)
- Tutti i metodi dichiarati in un'interfaccia sono implicitamente `public abstract`
- Tutti i campi dichiarati in un'interfaccia sono implicitamente `public static final`

### Implementare un'interfaccia - il contratto

- Implementare un'interfaccia equivale a firmare un contratto con il compilatore che stabilisce l'impegno ad implementare tutti i metodi specificati dall'interfaccia o a dichiarare la classe abstract
- 3 possibilità per una classe che implementa un'interfaccia:
  - fornire un'implementazione concreta di tutti i metodi, definendone il corpo
  - fornire l'implementazione concreta solo per un sottoinsieme proprio dei metodi dell'interfaccia
  - decidere di non fornire alcuna implementazione concreta
  - **Negli ultimi due casi, però, la classe va dichiarata `abstract`**

### Interfacce vs classi astratte

- Una domanda nasce spontanea: Se implementando un'interfaccia devo dichiarare tutti i metodi in essa definiti, perché non ricorrere ad una classe astratta?

#### Il problema del diamante

- Con l'ereditarietà multipla si possono creare situazione poco chiare di duplicazione di metodi e campi
- In Java non è consentito estendere più di una classe alla volta
- Al contrario, una classe può implementare tutte le interfacce desiderate
- È inoltre possibile estendere 1 classe e contemporaneamente implementare n interfacce diverse

#### Relazione tra intefacce e classi che le implementano

- Nel momento in cui una classe C decide di implementare un'interfacce I, tra queste due classi si instaura una relazione di tipo **is-a**, ovvero C è di tipo I
- Comportamento simile a quello dell'ereditarietà
- Quindi anche per le interfacce valgono le regole del polimorfismo

### Le classe interne

#### Classi top-level

- Le classi usate finora vengono dette top-level, cioè si trovano in alto di tutte le classi e non sono contenute in altre classi
- Questo tipo di classi richiedono un file `.java` dedicato con lo stesso nome della classe che esso contiene

#### Classi annidate (nested class)

- Java consente di scrivere classi all'interno di altre classi
- Le classi presenti all'interno sono chiamate classi annidate (nested class)
- Possono essere `static`
- Oppure `non-static`: in questo caso vengono dette classi interne (inner class)

#### Classi interne (inner class)

- Prima di poter creare un oggetto della classe interna è necessario istanziare la classe esterna che la contiene
- Ciascuna classe interna, infatti, ha un riferimento implicito all'oggetto della classe che lo contiene
- Dalla classe interna è possibile accedere a tutte le variabili e a tutti i metodi della classe esterna
- Come tutti i membri di una classe, le classi interne possono essere dichiarate `public`, `protected` o `private`
- Se dalla classe interna viene usato solamente `this` si fa riferimento ai campi e metodi di quella classe
- Per far riferimento alla classe esterna è necessario far procedere `this` dal nome della classe esterna e il punto

#### Classi annidate statiche (static nested class)

- Se invece una classe interna è statica, allora essa non richiede l'esistenza di un oggetto appartenente alla classe esterna e non ha nemmeno un riferimento implicito ad essa
  - Come i metodi statici, non può accedere allo stato dei singoli oggetti della classe esterna
- Da un punto di vista di comportamento, una classe annidata statica è equivalente ad una classe top-level inserita all'interno di un'altra classe top-level
- Sono accessibili all'esterno tramite il nome della classe esterna che le contiene, secondo la forma `new ClasseEsterna.ClasseInnestataStatica()`

#### Classi annidate: perché sono utili?

- Raggruppamento logico delle classi
  - Se una classe è utile solo a un'altra classe, è logico inserirla al suo interno e tenere le due classi logicamente vicine
- Incrementa l'incapsulamento
  - Una classe B annidata in A può accedere ai membri di A (anche se privati), ma può essere nascosta al mondo esterno
- Codice più leggibile e più facile da mantenere
  - La vicinanza spaziale è un fattore decisivo

## Strutture dati

### A cosa servono le strutture dati?

- A memorizzare e organizzare i dati in memoria così da utilizzarli in modo efficiente

### Caratteristiche di una struttura dati

- È necessario mantenere un **ordine**?
- Gli oggetti nella struttura possono **ripetersi**?
- È utile/necessario possedere una **chiave** per accedere a uno specifico oggetto?

### Le Collection

- Le collezioni in Java sono rese disponibili mediante il Java Collection Framework
- Strutture dati già pronte all'uso
  - con interfacce e algoritmi per modificarle
- Contengono e strutturano riferimenti ad altri oggetti
  - Tipicamente tutti dello stesso tipo
- Alcune interfacce del framework:
  - Collection: L'interfaccia alla radice della gerarchia delle collezioni
  - Set: Una collezione senza duplicati
  - List: Una collezione ordinata che può contenere duplicati
  - Map: Associa coppie di (**chiave, valore**), senza chiavi duplicate
  - Queue: Una collezione **first-in, first-out** che modella una coda

### Iterazione su una collezione

- Mediante gli Iterator:

  ```java
  // "vecchio" stile ma maggiore controllo
  Iterator<Integer> i = collezione.iterator();
  while(i.hasNext()) {
    int k = i.next();
    System.out.println(k);
  }
  ```

- Mediante il costrutto "for each":

  ```java
  // "nuovo" stile, più elegante
  for (Integer k : collezione) {
    System.out.println(k);
  }
  ```

- Mediante indici (solo su liste):

  ```java
  // su collezioni che implementano List: elemento per elemento
  for (int j = 0; j < collezione.size(); j++) {
    int k = collezione.get(j);
    System.out.println(k);
  }
  ```

### L'interfaccia java.util.Iterator<E>

- È un'interfaccia particolare che consente di iterare su collezioni
- Espone 3 metodi:
  - `boolean hasNext()`: Restituisce `true` se esiste ancora un successivo elemento nella collezione
  - `E next()`: Restituisce l'elemento successivo
  - `void remove()`: Rimuove l'elemento corrente (operazione opzionale)
- È in relazione con l'interfaccia `Iterable`, nel senso che chi implementa `Iterable` restituisce un `Iterator` sull'oggetto-collezione

### Collezioni fondamentali

- ArrayList
- LinkedList
  - Estendono AbstractList
  - Implementano List
- HashSet
- TreeSet
  - Estendono AbstractSet
  - Implementano Set
- HashMap
- TreeMap
  - Estendono AbstractMap
- LinkedHashMap
  - Estende HashMap

### Le liste: ArrayList e LinkedList

- Basate su List, una sottointerfaccia di Collection e Iterable
- Estendono AbstractList e implementano l'interfaccia List
- ArrayList implementa una lista mediante un array (eventualmente ridimensionato)
  - Capacità iniziale di 10 elementi
- LinkedList implementa una lista mediante elementi linkati

### Iterare sulle liste in entrambe le direzioni

- Utilizzando il metodo listIterator() restituito dall'interfaccia `List<E>`

```java
ArrayList<Integer> array = new ArrayList<>();
// specifico la posizione da cui partire
ListIterator<Integer> l = array.listIterator(array.size());
while (l.hasPrevious()) {
    System.out.println(l.previous());
}

```

### Insiemi: HashSet e TreeSet

- Gli insiemi sono Collection che contengono elementi tutti distinti
- HashSet memorizza gli elementi in una tabella di hash
- TreeSet memorizza gli elementi in un albero mantenendo un ordine sugli elementi

### Mappe

- Una mappa mette in corrispondenza chiavi e valori
- **Non** può contenere chiavi duplicate
- java.util.Map è un'interfaccia implementata da HashMap e da TreeMap
- HashMap memorizza le coppie in una tabella di hash
- TreeMap memorizza le coppie in un albero mantenendo l'ordine sulle chiavi
- LinkedHashMap estende HashMap e mantiene l'ordinamento di iterazione secondo gli inserimenti

> Es.
>
> Registriamo le frequenze di occorrenza delle parole in un testo utilizzando HashMap, TreeMap e LinkedHashMap.
>
> - Dato il seguente file: `a a e b c d e e b`
> - Se iteriamo su ciascuna delle tre strutture otteniamo:
>   - HashMap: {d=1, e=3, b=2, c=1, a=2} - Senza ordinamento
>   - TreeMap: {a=2, b=2, c=1, d=1, e=3} - Ordinamento naturale (String)
>   - LinkedHashMap: {a=2, e=3, b=2, c=1, d=1} - Ordinamento di inserimento

#### Chiavi e valori di una mappa

- È possibile ottenere l'insieme delle chiavi di una mappa mediante il metodo `keySet`
- L'elenco dei valori mediante il metodo `values` (con ripetizione)
- L'insieme delle coppie (chiave, valore) mediante il metodo `entrySet`

### Pila e Coda

- Due strutture dati fondamentali, utili in un gran numero di attività
- Coda (FIFO: First In, First Out)
  - Come una coda di persone alle poste, il primo a essere ricevuto è il primo a uscire
- Pila o stack (LIFO: Last In, First Out)
  - Come una pila di piatti, l'ultimo impilato sarà il primo a uscire

#### Coda

- Esempi di coda:
  - Coda degli eventi: relativi a mouse e tastiera
  - Coda di stampa
- Esistono implementazioni standard della coda mediante l'interfaccia Queue
  - LinkedList implementa l'interfaccia Queue
- Operazioni principali:
  - `add`: inserisce un elemento in coda
  - `remove`: rimuove un elemento all'inizio della coda
  - `peek`: restituisce l'elemento all'inizio della coda senza rimuoverlo

#### Pila

- Esempi di pila:
  - La pila di esecuzione (run-time stack) contenente i record di attivazione delle chiamate a metodi
  - Nell'implementazione della ricorsione
- Esiste un'implementazione standard mediante la classe Stack
  - Implementa l'interfaccia List
- Operazioni principali
  - `push`: inserisce un elemento in cima alla pila
  - `pop`: rimuove l'elemento in cima alla pila
  - `peek`: restituisce l'elemento in cima alla pila senza rimuoverlo

### Come scegliere una Collection

- Come voglio accedere agli elementi?
  - Tramite una posizione? ArrayList
  - Tramite una chiave? mappa
  - Non importa: un insieme
- Quali sono i tipi degli elementi o tipi di chiavi e valori?
  - Ad esempio, in una mappa, vado da stringhe a interi o da interi a stringhe? Qual è la chiave univoca che deve determinare il valore?
- L'ordine degli elementi o delle chiavi è importante?
  - Sì, in un ordine naturale degli elementi: TreeMap o TreeSet
  - Sì, nell'ordine dell'inserimento: ArrayList, LinkedList o LinkedHashSet
  - Non importa: HashSet o HashMap
- Per una collezione non mappa, quali operazioni devono essere veloci?
  - Ricerca veloce: HashSet
  - Aggiunta e rimozione di elementi: LinkedList
  - Non importa/ho pochi elementi nella collezione: ArrayList
- Se usate rappresentazioni ad albero (TreeSet o TreeMap), serve che si implementi Comparable?
  - No, se l'ordine naturale è già prestabilito (es. tipi primitivi)
  - Sì, è un tipo nuovo per il quale vogliamo stabilire l'ordinamento

### Gli Alberi

- Una struttura dati ricorsiva in cui ogni nodo possiede un padre tranne la radice
- Gli alberi più comuni sono binari (ovvero con al più due figli per nodo)
