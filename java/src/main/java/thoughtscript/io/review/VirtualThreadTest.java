package thoughtscript.io.review;

// Testing claims made here:
// https://blog.rockthejvm.com/ultimate-guide-to-java-virtual-threads/#10-appendix-maven-configuration
// https://davidvlijmincx.com/posts/use-executorservice-with-virtual-threads/

import lombok.SneakyThrows;

import java.lang.management.ManagementFactory;
import java.util.concurrent.*;

public class VirtualThreadTest {
    private static Executor executor = Executors.newVirtualThreadPerTaskExecutor();

    private static void printResourceInfo(String id) {
        System.out.println("====================================================================================================");

        System.out.println(id);

        System.out.println("Number of Thread.activeCount() threads " + Thread.activeCount());
        System.out.println("Total Number of ManagementFactory.getThreadMXBean().getThreadCount() threads " + ManagementFactory.getThreadMXBean().getThreadCount());

        long heapSize = Runtime.getRuntime().totalMemory();
        long heapMaxSize = Runtime.getRuntime().maxMemory();
        long heapFreeSize = Runtime.getRuntime().freeMemory();
        System.out.println("heapSize: " + heapSize + " heapMaxSize: " + heapMaxSize + " heapFreeSize: "+ heapFreeSize);

        System.out.println("====================================================================================================");
    }

    @SneakyThrows // throws ExecutionException, InterruptedException

    // https://projectlombok.org/features/SneakyThrows
    // an implicit effectively wrapping try-catch clause that doesn't modify function signature
    public static void testThreadOfVirtual() {
        Thread thread = Thread.ofVirtual().name("thread-1").start(() -> {
            String id = "threadOfVirtual() > start()";
            printResourceInfo(id);
            System.out.println("Completed: " + id);
        });

        thread.join();
    }

    @SneakyThrows // throws ExecutionException, InterruptedException
    public static void testNewVirtualThreadPerTaskExecutor() {
        // Returns a Future
        Future<String> f = ((ExecutorService) executor).submit(()-> {
            String result = "newVirtualThreadPerTaskExecutor() > submit()";
            printResourceInfo(result);
            return result;
        });

        System.out.println("Completed: " + f.get());
    }
}
