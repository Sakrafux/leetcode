import java.util.concurrent.Semaphore;

public class Main {

    public static void main(String[] args) {
        runCase(1);
        runCase(2);
    }

    private static void runCase(int n) {
        var foobar = new FooBar(n);
        var thread1 = new Thread(() -> {
            try {
                foobar.foo(() -> System.out.print("foo"));
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        var thread2 = new Thread(() -> {
            try {
                foobar.bar(() -> System.out.print("bar"));
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        thread1.start();
        thread2.start();
        try {
            thread1.join();
            thread2.join();
        } catch (InterruptedException ignored) {}
        System.out.println();
    }

}

class FooBar {
    private int n;

    private Semaphore semA = new Semaphore(1);
    private Semaphore semB = new Semaphore(0);

    public FooBar(int n) {
        this.n = n;
    }

    public void foo(Runnable printFoo) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            semA.acquire();
            // printFoo.run() outputs "foo". Do not change or remove this line.
            printFoo.run();
            semB.release();
        }
    }

    public void bar(Runnable printBar) throws InterruptedException {
        for (int i = 0; i < n; i++) {
            semB.acquire();
            // printBar.run() outputs "bar". Do not change or remove this line.
            printBar.run();
            semA.release();
        }
    }
}