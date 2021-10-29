代码地址 ： https://github.com/skyhackvip/dispatcher


每个节点设置对应的权重，权重越大可能被选中的次数越高，某节点被选中的次数≈(本节点权重/全部权重) * 总分配次数。
举例说明：
3个节点对应权重为 node:weight [a=1,b=2,c=5]
currentWeight 代表每次请求节点的当前权重，为currentWeight+weight
totalWeight  代表所有节点初始权重之和 1+2+5=8
第一次请求：
currentWeight 为 [a=0+1,b=0+2,c=0+5] ，选中最大的c做为本次输出，之后c节点的权重需要减去totalWeightcurrent，调整后 [a=1,b=2,c=5-8]
也就是 [a=1,b=2,c=-3]
第二次请求：
currentWeight 为  [a=1+1,b=2+2,c=-3+5]  结果为 [a=2,b=4,c=2]，选中最大的b作为本次输出，之后节点权重变更为 [a=2,b=-4,c=2]
第三次请求：
currentWeight 为  [a=2+1,b=-4+2,c=2+5]  结果为 [a=3,b=-2,c=7]，又轮到c（权重大的好处体现出来了），之后节点权重变更为 [a=3,b=-2,c=-1]
第四次请求：
 [a=3,b=-2,c=-1] 加权后[a=4,b=0,c=4]，a与c相等，优先选前者输出a