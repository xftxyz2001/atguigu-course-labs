import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

@Retention(RetentionPolicy.RUNTIME)
@interface JSONField {
    String name() default "";
}

// 定义了一个Monster类
class Monster {

    @JSONField(name = "name")
    String name;

    @JSONField(name = "monster_age")
    int age;

    @JSONField(name = "成绩")
    float score;

    String sex;

    // 方法，返回两个数的和
    int getSum(int n1, int n2) {
        return n1 + n2;
    }

    // 方法，接收四个值，给Monster赋值
    void set(String name, int age, float score, String sex) {
        this.name = name;
        this.age = age;
        this.score = score;
        this.sex = sex;
    }

    // 方法，显示s的值
    void print() {
        System.out.println("---start~----");
        System.out.println(this);
        System.out.println("---end~----");
    }

    @Override
    public String toString() {
        return "{" + this.name + " " + this.age + " " + this.score + " " + this.sex + "}";
    }
}

class Main {

    static void testStruct(Object obj)
            throws NoSuchMethodException, SecurityException, IllegalAccessException, InvocationTargetException {
        // 获取类型
        Class c = obj.getClass();

        // 获取所有属性
        Field[] fields = c.getDeclaredFields();
        System.out.println("Class has " + fields.length + " fields");
        for (Field field : fields) {
            // 打印属性值
            System.out.println(field.getName() + ": " + field.get(obj));

            // 获取属性上的注解
            JSONField annotation = field.getDeclaredAnnotation(JSONField.class);
            if (annotation != null) {
                System.out.println(field.getName() + " has annotation: " + annotation.name());
            }
        }

        // 获取所有方法
        Method[] methods = c.getDeclaredMethods();
        System.out.println("Class has " + methods.length + " methods");
        // 调用print方法
        Method print = c.getDeclaredMethod("print");
        print.invoke(obj);

        // 调用getSum方法
        Method getSum = c.getDeclaredMethod("getSum", int.class, int.class);
        Object res = getSum.invoke(obj, 10, 40);
        System.out.println("res=" + res);

    }

    public static void main(String[] args) throws NoSuchMethodException, SecurityException, IllegalAccessException,
            IllegalArgumentException, InvocationTargetException {
        Monster monster = new Monster();
        monster.set("黄鼠狼精", 400, 30.8f, null);
        testStruct(monster);
    }

}

// Class has 4 fields
// name: 黄鼠狼精
// name has annotation: name
// age: 400
// age has annotation: monster_age
// score: 30.8
// score has annotation: 成绩
// sex: null
// Class has 4 methods
// ---start~----
// {黄鼠狼精 400 30.8 null}
// ---end~----
// res=50
