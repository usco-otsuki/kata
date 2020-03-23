using System.Collections.Generic;
using System;
using System.Linq;

namespace StaticArithmeticMachine
{
    public class Cpu
    {
        private Dictionary<string, int> registers;
        private Stack<int> stack;
        public Cpu()
        {
            this.registers = new Dictionary<string, int> { { "a", 0 }, { "b", 0 }, { "c", 0 }, { "d", 0 } };
            this.stack = new Stack<int>();
        }
        // Returns the value of the named register.
        public int ReadReg(string name)
        {
            return registers[name];
        }
        // Stores the value into the given register.
        public void WriteReg(string name, int value)
        {
            registers[name] = value;
        }

        // Pops the top element of the stack, returning the value.
        public int PopStack()
        {
            return stack.Pop();
        }
        //  Pushes an element onto the stack.
        public void WriteStack(int value)
        {
            stack.Push(value);
        }
        // Prints the contents of the stack to System.Console.
        public void PrintStack()
        {
            foreach (int val in stack)
            {
                Console.WriteLine(val);
            }
        }

        public class Machine
        {
            private Cpu cpu;

            public Machine(Cpu cpu = null)
            {
                this.cpu = cpu;
            }

            private bool isNonNegativeInteger(string str)
            {
                int num;
                bool parseable = int.TryParse(str, out num);
                if (!parseable)
                {
                    return false;
                }
                return num >= 0;
            }

            private bool isRegister(string str)
            {
                return new[] { "a", "b", "c", "d" }.Contains(str);
            }

            public void Exec(string op)
            {
                var tokens = op.Split(new char[] { ' ' });
                string cmd = tokens[0];
                string[] args = tokens.Skip(1).ToArray();
                switch (cmd)
                {
                    case "push":
                        if (isNonNegativeInteger(args[0]))
                        {
                            int val = int.Parse(args[0]);
                            cpu.WriteStack(val);
                            return;
                        }
                        else if (isRegister(args[0]))
                        {
                            cpu.WriteStack(cpu.ReadReg(args[0]));
                        }
                        else
                        {
                            throw new ArgumentException();
                        }
                        return;
                    case "pop":
                        {
                            int val = cpu.PopStack();
                            if (args.Length > 0)
                            {
                                cpu.WriteReg(args[0], val);
                            }
                            return;
                        }
                    case "pushr":
                        cpu.WriteStack(cpu.ReadReg("a"));
                        cpu.WriteStack(cpu.ReadReg("b"));
                        cpu.WriteStack(cpu.ReadReg("c"));
                        cpu.WriteStack(cpu.ReadReg("d"));
                        return;
                    case "pushrr":
                        cpu.WriteStack(cpu.ReadReg("d"));
                        cpu.WriteStack(cpu.ReadReg("c"));
                        cpu.WriteStack(cpu.ReadReg("b"));
                        cpu.WriteStack(cpu.ReadReg("a"));
                        return;
                    case "popr":
                        cpu.WriteReg("d", cpu.PopStack());
                        cpu.WriteReg("c", cpu.PopStack());
                        cpu.WriteReg("b", cpu.PopStack());
                        cpu.WriteReg("a", cpu.PopStack());
                        return;
                    case "poprr":
                        cpu.WriteReg("a", cpu.PopStack());
                        cpu.WriteReg("b", cpu.PopStack());
                        cpu.WriteReg("c", cpu.PopStack());
                        cpu.WriteReg("d", cpu.PopStack());
                        return;
                    case "mov":
                        if (isNonNegativeInteger(args[0]))
                        {
                            int val = int.Parse(args[0]);
                            cpu.WriteReg(args[1], val);
                            return;
                        }
                        else if (isRegister(args[0]))
                        {
                            cpu.WriteReg(args[0], 0); // is this needed?
                            cpu.WriteReg(args[1], cpu.ReadReg(args[0]));
                        }
                        else
                        {
                            throw new ArgumentException();
                        }
                        return;
                    case "add":
                    case "sub":
                    case "mul":
                    case "div":
                    case "and":
                    case "or":
                    case "xor":
                        int calcResult = 0;
                        if (args.Length == 0)
                        {
                            throw new ArgumentException();
                        }
                        var operands = new List<int>();
                        if (isNonNegativeInteger(args[0]) || isRegister(args[0]))
                        {
                            int val;
                            if (isNonNegativeInteger(args[0]))
                            {
                                val = int.Parse(args[0]);
                            }
                            else
                            {
                                val = cpu.ReadReg(args[0]);
                            }
                            for (int i = 0; i < val; i++)
                            {
                                operands.Add(cpu.PopStack());
                            }
                        }
                        else
                        {
                            throw new ArgumentException();
                        }
                        switch (cmd)
                        {
                            case "add":
                                calcResult = operands.Aggregate(0, (acc, x) => acc + x);
                                break;
                            case "sub":
                                calcResult = operands.Aggregate(0, (acc, x) => acc - x);
                                break;
                            case "mul":
                                calcResult = operands.Aggregate(1, (acc, x) => acc * x);
                                break;
                            case "div":
                                calcResult = operands.Skip(1).Aggregate(operands.First(), (acc, x) => acc / x);
                                break;
                            case "and":
                                calcResult = operands.Aggregate(1, (acc, x) => acc & x);
                                break;
                            case "or":
                                calcResult = operands.Aggregate(0, (acc, x) => acc | x);
                                break;
                            case "xor":
                                calcResult = operands.Skip(1).Aggregate(operands.First(), (acc, x) => acc ^ x);
                                break;
                        }
                        if (args.Length == 1)
                        {
                            cpu.WriteReg("a", calcResult);
                            return;
                        }
                        else if (args.Length == 2 && isRegister(args[1]))
                        {
                            cpu.WriteReg(args[1], calcResult);
                            return;
                        }
                        throw new ArgumentException();
                }

            }
        }
    }
}