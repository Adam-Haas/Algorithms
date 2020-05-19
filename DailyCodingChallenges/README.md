[![Daily Interview Pro](http://sea.dailyinterviewpro.com/logo-small.png)](https://nam03.safelinks.protection.outlook.com/?url=http%3A%2F%2Fwww.dailyinterviewpro.com%2F&data=02%7C01%7C%7Cdf730f7948ba4bfe5a2408d7fc060369%7C84df9e7fe9f640afb435aaaaaaaaaaaa%7C1%7C0%7C637254975062038972&sdata=WAQLet2jjyb9xyw0vrJcAiOR7TI2%2F5AXDG2%2F8WWsfW4%3D&reserved=0)

Hi, here's your problem today. This problem was recently asked by Microsoft:  
  
You are given two linked-lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.  
  
Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)  
Output: 7 -> 0 -> 8  
Explanation: 342 + 465 = 807.  

Here is the function signature as a starting point (in Python):  
  

\# Definition for singly-linked list.  
class ListNode(object):  
 def \_\_init\_\_(self, x):  
 self.val \= x  
 self.next \= None  
  
class Solution:  
 def addTwoNumbers(self, l1, l2, c \= 0):  
 \# Fill this in.  
  
l1 \= ListNode(2)  
l1.next \= ListNode(4)  
l1.next.next \= ListNode(3)  
  
l2 \= ListNode(5)  
l2.next \= ListNode(6)  
l2.next.next \= ListNode(4)  
  
result \= Solution().addTwoNumbers(l1, l2)  
while result:  
 print result.val,  
 result \= result.next  
\# 7 0 8  

  
**Why Python?** We recommend using Python as a generalist language for interviewing, as it is well-regarded in the tech industry and used across Google/YouTube, Facebook/Instagram, Netflix, Uber, Dropbox, Pinterest, Spotify, etc., It is easy to learn with readable syntax, and very similar in structure to other popular languages like Java, C/C++, Javascript, PHP, Ruby, etc. Python is generally faster to read/write though, which makes it ideal for interviews. You can, of course, use any language you like!  
  
[Upgrade to PRO](https://nam03.safelinks.protection.outlook.com/?url=https%3A%2F%2Fwww.techseries.dev%2Foffers%2Fvb9SVFHL%2Fcheckout&data=02%7C01%7C%7Cdf730f7948ba4bfe5a2408d7fc060369%7C84df9e7fe9f640afb435aaaaaaaaaaaa%7C1%7C0%7C637254975062038972&sdata=d5931L%2FkYD6bHCfUqSpgKhn09s4Ug6X4Y3TJddR%2BmLI%3D&reserved=0) and get in-depth solutions to every problem from ex-Google and ex-Facebook engineers TechLead and Joma.  
  
Â» Ready to fast-track your career? Join our premiere tech interview training program [Tech Interview Pro](https://nam03.safelinks.protection.outlook.com/?url=https%3A%2F%2Fwww.techseries.dev%2Finterview-details&data=02%7C01%7C%7Cdf730f7948ba4bfe5a2408d7fc060369%7C84df9e7fe9f640afb435aaaaaaaaaaaa%7C1%7C0%7C637254975062043953&sdata=m6DYSoodfF0OG9VGpjmm1WrigX0selkTrnp9sbBCONk%3D&reserved=0).  
Â» And join us at CoderPro for 100+ whiteboard coding sessions: [https://coderpro.com/](https://nam03.safelinks.protection.outlook.com/?url=https%3A%2F%2Fcoderpro.com%2F&data=02%7C01%7C%7Cdf730f7948ba4bfe5a2408d7fc060369%7C84df9e7fe9f640afb435aaaaaaaaaaaa%7C1%7C0%7C637254975062043953&sdata=qIiLH7DyGoU6ynYvQejQIAiRFBFNWDzOaYkiQunbfvE%3D&reserved=0) (20% off now, limited time)  
  
Have a great day!  
Daily Interview Pro

© 2020 Daily Interview Pro. All rights reserved.
