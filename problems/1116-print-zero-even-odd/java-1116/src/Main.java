import java.util.concurrent.Semaphore;
import java.util.function.IntConsumer;

public class Main {

    public static void main(String[] args) {
        runCase(2);
        runCase(5);
    }

    private static void runCase(int n) {
        IntConsumer consumer = System.out::print;
        var zeo = new ZeroEvenOdd(n);
        var thread1 = new Thread(() -> {
            try {
                zeo.zero(consumer);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        var thread2 = new Thread(() -> {
            try {
                zeo.even(consumer);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        var thread3 = new Thread(() -> {
            try {
                zeo.odd(consumer);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        thread1.start();
        thread2.start();
        thread3.start();
        try {
            thread1.join();
            thread2.join();
            thread3.join();
        } catch (InterruptedException ignored) {}
        System.out.println();
    }

}

class ZeroEvenOdd {
    private int n;

    private boolean isOdd = true;

    private Semaphore semZero = new Semaphore(1);
    private Semaphore semEven = new Semaphore(0);
    private Semaphore semOdd = new Semaphore(0);

    public ZeroEvenOdd(int n) {
        this.n = n;
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            semZero.acquire();
            printNumber.accept(0);
            if (isOdd) {
                isOdd = false;
                semOdd.release();
            } else {
                isOdd = true;
                semEven.release();
            }
        }
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i = 2; i <= n; i += 2) {
            semEven.acquire();
            printNumber.accept(i);
            semZero.release();
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i = 1; i <= n; i += 2) {
            semOdd.acquire();
            printNumber.accept(i);
            semZero.release();
        }
    }
}