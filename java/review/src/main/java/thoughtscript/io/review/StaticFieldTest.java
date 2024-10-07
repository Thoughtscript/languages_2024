package thoughtscript.io.review;

public class StaticFieldTest {
    public static String testfield = "I'm a static field";

    public void forceExposeStaticFieldThroughGetter() {
        // System.out.println(this.testfield); 
        // Cannot mix and match static and non-static still - will through exception

        System.out.println("Cannot mix and match static and non-static still!");
    }
}
