public class Main {

    public static void main(String[] args) {
        var foo = new Foo();
        new Thread(() -> {
            try {
                foo.first(() -> System.out.print("first"));
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }).start();
        new Thread(() -> {
            try {
                foo.second(() -> System.out.print("second"));
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }).start();
        new Thread(() -> {
            try {
                foo.third(() -> System.out.print("third"));
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        }).start();
    }

}

class Foo {

    private int state = 0;

    public Foo() {

    }

    public synchronized void first(Runnable printFirst) throws InterruptedException {

        // printFirst.run() outputs "first". Do not change or remove this line.
        printFirst.run();

        state = 1;
        this.notifyAll();
    }

    public synchronized void second(Runnable printSecond) throws InterruptedException {
        while (state != 1) {
            this.wait();
        }

        // printSecond.run() outputs "second". Do not change or remove this line.
        printSecond.run();

        state = 2;
        this.notifyAll();
    }

    public synchronized void third(Runnable printThird) throws InterruptedException {
        while (state != 2) {
            this.wait();
        }

        // printThird.run() outputs "third". Do not change or remove this line.
        printThird.run();

        this.notifyAll();
    }

}