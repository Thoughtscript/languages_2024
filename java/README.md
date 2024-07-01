# Java Review Topics 2024

[![](https://img.shields.io/badge/Maven-3.9.6-white.svg)](https://maven.apache.org/download.cgi)
[![](https://img.shields.io/badge/Docker-blue.svg)](https://www.docker.com/)
[![](https://img.shields.io/badge/Java-1.22-blue.svg)](https://www.oracle.com/java/technologies/downloads/#java22)

Test client.

> Helps avoid local Java SDK Version, `pom.xml`, and Maven mismatches:

```bash
docker build .
docker images # get image id
docker run <IMAGE_ID>
```

## Virtual Threads

Looking into Java 21+ lightweight Virtual Threads.

> Baeldung (whom I used to write for and which is Top 3 most-trusted Java resource online!) suggests that JVT's can replace WebFlux in some scenarios.

### Configuration

```
-Djdk.virtualThreadScheduler.parallelism=2
-Djdk.virtualThreadScheduler.maxPoolSize=2
-Djdk.virtualThreadScheduler.minRunnable=2
```

### Assessing the Accuracy of Articles

1. https://blog.rockthejvm.com/ultimate-guide-to-java-virtual-threads/ <- This is incorrect 

   "Also, they were designed with the idea of using a different virtual thread for each request. So, it’s worthless to use a thread pool or an executor service to create virtual threads."

2. From the official documentation: 
   * https://docs.oracle.com/en%2Fjava%2Fjavase%2F22%2Fdocs%2Fapi%2F%2F/java.base/java/util/concurrent/Executors.html#newVirtualThreadPerTaskExecutor() 
   * https://docs.oracle.com/en/java/javase/21/core/virtual-threads.html#GUID-2DDA5807-5BD5-4ABC-B62A-A1230F0566E0

Looks like:

1. heapFreeSize in `newVirtualThreadPerTaskExecutor()` <= `threadOfVirtual()` even when `threadOfVirtual()` is executed first and in the same `static` context.
2. Interestingly enough `Thread.activeCount()` will not report `Virtual Threads` (even `newVirtualThreadPerTaskExecutor()`)!
3. Unlike other `ServiceExecutors`, `newVirtualThreadPerTaskExecutor()` is [unbounded](https://docs.oracle.com/en%2Fjava%2Fjavase%2F22%2Fdocs%2Fapi%2F%2F/java.base/java/util/concurrent/Executors.html#newVirtualThreadPerTaskExecutor()) in terms of how many can be launched.

```bash
====================================================================================================
threadOfVirtual() > start()
Number of Thread.activeCount() threads 0
Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads 9
heapSize: 264241152 heapMaxSize: 4185915392 heapFreeSize: 260053584
====================================================================================================
Completed: threadOfVirtual() > start()
====================================================================================================
newVirtualThreadPerTaskExecutor() > submit()
Number of Thread.activeCount() threads 0
Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads 9
heapSize: 264241152 heapMaxSize: 4185915392 heapFreeSize: 259550216
====================================================================================================
Completed: newVirtualThreadPerTaskExecutor() > submit()
====================================================================================================
newVirtualThreadPerTaskExecutor() > submit()
Number of Thread.activeCount() threads 0
Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads 9
heapSize: 264241152 heapMaxSize: 4185915392 heapFreeSize: 259550216
====================================================================================================
Completed: newVirtualThreadPerTaskExecutor() > submit()
====================================================================================================
threadOfVirtual() > start()
Number of Thread.activeCount() threads 0
Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads 9
heapSize: 264241152 heapMaxSize: 4185915392 heapFreeSize: 259550216
====================================================================================================
Completed: threadOfVirtual() > start()
```
```bash
====================================================================================================
threadOfVirtual() > start()
Number of Thread.activeCount() threads 0
Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads 9
heapSize: 264241152 heapMaxSize: 4185915392 heapFreeSize: 260053584
====================================================================================================
Completed: threadOfVirtual() > start()
====================================================================================================
newVirtualThreadPerTaskExecutor() > submit()
Number of Thread.activeCount() threads 0
Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads 9
heapSize: 264241152 heapMaxSize: 4185915392 heapFreeSize: 259550216
====================================================================================================
Completed: newVirtualThreadPerTaskExecutor() > submit()
====================================================================================================
newVirtualThreadPerTaskExecutor() > submit()
Number of Thread.activeCount() threads 0
Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads 9
heapSize: 264241152 heapMaxSize: 4185915392 heapFreeSize: 259550216
====================================================================================================
Completed: newVirtualThreadPerTaskExecutor() > submit()
====================================================================================================
threadOfVirtual() > start()
Number of Thread.activeCount() threads 0
Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads 9
heapSize: 264241152 heapMaxSize: 4185915392 heapFreeSize: 259550216
====================================================================================================
Completed: threadOfVirtual() > start()
```

## Resources and Links

1. https://www.baeldung.com/java-get-number-of-threads
2. https://www.baeldung.com/java-reactor-webflux-vs-virtual-threads
3. https://projectlombok.org/features/SneakyThrows
4. https://digma.ai/a-list-of-major-java-and-jvm-features-since-jdk-17-to-22/