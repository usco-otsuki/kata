using Xunit;

namespace StackArithmeticMachine
{

    public class SolutionTest
    {

        private Cpu cpu;
        private Machine machine;

        public SolutionTest()
        {
            this.cpu = new Cpu();
            this.machine = new Machine(cpu);
        }

        [Fact]
        public void AdditionTest()
        {
            machine.Exec("push 25");
            machine.Exec("push 15");
            machine.Exec("push 10");
            machine.Exec("add 3");

            Assert.Equal(50, cpu.ReadReg("a"));
        }

        [Fact]
        public void SubtractionTest()
        {
            machine.Exec("push 22");
            machine.Exec("push 13");
            machine.Exec("push 75");
            machine.Exec("sub 3");

            Assert.Equal(40, cpu.ReadReg("a"));
        }

        [Fact]
        public void MultiplicationTest()
        {
            machine.Exec("push 10");
            machine.Exec("push 5");
            machine.Exec("push 5");
            machine.Exec("push 5");
            machine.Exec("push 5");
            machine.Exec("mul 5");

            Assert.Equal(6250, cpu.ReadReg("a"));
        }

        [Fact]
        public void DivisionTest()
        {
            machine.Exec("push 5");
            machine.Exec("push 20");
            machine.Exec("push 5");
            machine.Exec("push 10000");
            machine.Exec("div 4");

            Assert.Equal(20, cpu.ReadReg("a"));
        }

        [Fact]
        public void AndTest()
        {
            machine.Exec("push 7");
            machine.Exec("push 5");
            machine.Exec("push 4");
            machine.Exec("and 3, c");

            Assert.Equal(4, cpu.ReadReg("c"));
        }


        [Fact]
        public void OrTest()
        {
            machine.Exec("push 0");
            machine.Exec("push 2");
            machine.Exec("push 4");
            machine.Exec("or 3, d");

            Assert.Equal(6, cpu.ReadReg("d"));
        }

        [Fact]
        public void XorTest()
        {
            machine.Exec("push 2");
            machine.Exec("push 5");
            machine.Exec("push 4");
            machine.Exec("push 5");
            machine.Exec("push 2");
            machine.Exec("xor 5, b");

            Assert.Equal(4, cpu.ReadReg("b"));
        }

        [Fact]
        public void PopTest()
        {
            machine.Exec("push 1");
            machine.Exec("push 2");
            machine.Exec("pop");
            machine.Exec("pop b");

            Assert.Equal(1, cpu.ReadReg("b"));
        }

        [Fact]
        public void PushPopInOrderTest()
        {
            machine.Exec("mov 1 a");
            machine.Exec("mov 2 b");
            machine.Exec("mov 3 c");
            machine.Exec("mov 4 d");
            machine.Exec("pushr");
            machine.Exec("popr");
            Assert.Equal(1, cpu.ReadReg("a"));
            Assert.Equal(2, cpu.ReadReg("b"));
            Assert.Equal(3, cpu.ReadReg("c"));
            Assert.Equal(4, cpu.ReadReg("d"));
        }

        [Fact]
        public void PushPopReverseTest()
        {
            machine.Exec("mov 1 a");
            machine.Exec("mov 2 b");
            machine.Exec("mov 3 c");
            machine.Exec("mov 4 d");
            machine.Exec("pushrr");
            machine.Exec("poprr");
            Assert.Equal(1, cpu.ReadReg("a"));
            Assert.Equal(2, cpu.ReadReg("b"));
            Assert.Equal(3, cpu.ReadReg("c"));
            Assert.Equal(4, cpu.ReadReg("d"));
        }

        [Fact]
        public void PushInOrderPopReverseTest()
        {
            machine.Exec("mov 1 a");
            machine.Exec("mov 2 b");
            machine.Exec("mov 3 c");
            machine.Exec("mov 4 d");
            machine.Exec("pushr");
            machine.Exec("poprr");
            Assert.Equal(4, cpu.ReadReg("a"));
            Assert.Equal(3, cpu.ReadReg("b"));
            Assert.Equal(2, cpu.ReadReg("c"));
            Assert.Equal(1, cpu.ReadReg("d"));
        }


        [Fact]
        public void SimpleArithmeticTest()
        {
            machine.Exec("push 25");
            machine.Exec("push 15");
            machine.Exec("push 10");
            machine.Exec("add 3");

            machine.Exec("push 17");
            machine.Exec("suba 2, b");

            machine.Exec("push 5");
            machine.Exec("mula 2, c");

            machine.Exec("pushr");
            machine.Exec("add 4");

            Assert.Equal(333, cpu.ReadReg("a"));
        }


        [Fact]
        public void TriangleAngleTest()
        {
            machine.Exec("mov 15, a");
            machine.Exec("mov 22, b");
            machine.Exec("mov 2, c");

            machine.Exec("pushrr");
            machine.Exec("mul 2");
            machine.Exec("diva 2");

            Assert.Equal(165, cpu.ReadReg("a"));
        }

        [Fact]
        public void DiscardTest()
        {
            machine.Exec("push 25");
            machine.Exec("push 15");
            machine.Exec("push 10");
            machine.Exec("mov 0, a");
            machine.Exec("adda 2");
            machine.Exec("pop");
            machine.Exec("adda 2");

            Assert.Equal(35, cpu.ReadReg("a"));
        }
    }
}