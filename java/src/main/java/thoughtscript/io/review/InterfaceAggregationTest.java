package thoughtscript.io.review;

import lombok.NoArgsConstructor;

public class InterfaceAggregationTest {

    public interface A {

        void voidMethod();

        class D implements A {
            String x;

            public D (String x) {
                this.x = x;
            }

            @Override
            public void voidMethod() {
                System.out.println(this.x);
            }
        }

        class G extends B.E {
            public void test() {
                super.voidMethod();
            }
        }

        interface B extends A {

            @NoArgsConstructor
            class C implements B {

                private static String name;

                private static final String staticName = "staticName";

                private final D d = new D("y");

                public void setName(String name) {
                    this.name = name;
                }

                public String getName() {
                    return this.name;
                }

                public static void getStaticName() {
                    System.out.println(staticName);
                }

                @Override
                public void voidMethod() {
                    this.d.voidMethod();
                    setName(this.d.x);
                    System.out.println(getName());
                }
            }

            class E implements B {
                @Override
                public void voidMethod() {
                    System.out.println("e");
                }
            }

            class F extends C {

                public F () {
                    super.setName("f");
                }

                public void test() {
                    super.voidMethod();
                }
            }
        }
    }
}
