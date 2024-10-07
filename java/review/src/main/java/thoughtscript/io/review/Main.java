package thoughtscript.io.review;

public class Main {

    public static void main(String[] args) {
        InterfaceAggregationTest.A.B.D d = new InterfaceAggregationTest.A.D("A.B.D > d");
        d.voidMethod();

        InterfaceAggregationTest.A.B.F f = new InterfaceAggregationTest.A.B.F();
        f.test();

        VirtualThreadTest.testThreadOfVirtual();
        VirtualThreadTest.testNewVirtualThreadPerTaskExecutor();

        VirtualThreadTest.testNewVirtualThreadPerTaskExecutor();
        VirtualThreadTest.testThreadOfVirtual();

        VirtualThreadTest.testThreadOfVirtual();
        VirtualThreadTest.testNewVirtualThreadPerTaskExecutor();

        VirtualThreadTest.testNewVirtualThreadPerTaskExecutor();
        VirtualThreadTest.testThreadOfVirtual();

        StaticFieldTest staticFieldTest = new StaticFieldTest();
        staticFieldTest.forceExposeStaticFieldThroughGetter();
        System.out.println(StaticFieldTest.testfield);
    }
}