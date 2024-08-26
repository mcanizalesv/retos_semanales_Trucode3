# Interactions

## A Person

When you create a person, you need to give them a name. All new persons start with specific stats: 80 points of hunger, 10 points of sleepiness, and 80 points of stamina. Each of these stats should remain in the range of 0 to 100.

### Checking Status

If you ask a person about their status, they will respond with:

```
My name is {{name}}, this is my status:
Hunger: {{hunger}}
Sleepyness: {{sleepyness}}
Stamina: {{stamina}}
```

They will also provide status alerts based on their stats:

- "I am very hungry" if hunger is 100.
- "I am satiated" if hunger is 0.
- "I need to sleep" if sleepiness is 100.
- "I am totally rested" if sleepiness is 0.
- "I have a lot of energy" if stamina is 100.
- "I'm tired" if stamina is 0.

### Exercising

When you ask a person to exercise, you must specify the intensity level ("low," "medium," or "high"). Here's how it affects their stats:

- **low** intensity: "I have exercised at low intensity." Stamina decreases by 10 points, and hunger increases by 10 points.
- **medium** intensity: "I have exercised at medium intensity." Stamina decreases by 25 points, and hunger increases by 30 points.
- **high** intensity: "I have exercised at high intensity." Stamina decreases by 50 points, and hunger increases by 60 points.

### Sleeping

If you tell a person to sleep, their hunger increases by 20 points, sleepiness goes to zero, and stamina rises to 100. They will say, "I have slept."

### Studying

When you ask a person to study, their hunger increases by 25 points, sleepiness increases by 30 points, and stamina decreases by 10 points. They will respond with, "I have studied."

### Eating

If you instruct a person to eat, you must provide them with food. They will adjust their stats based on the food's effects. Then they will say, "I have eaten," and mention, "That {{food name}} tasted {{food taste}}!"

## Child (A Special Person)

Children are a special kind of person. They are created exactly like persons, but their hunger starts at 100 points. Additionally, when they eat food, they gain only 80% of the hunger effect, 110% of the sleepiness effect, and 120% of the stamina effect.

## Food (Interface for Various Types)

### Creating Food

When creating food, you must provide it with various fields:

- name
- taste
- effects
  - hunger
  - stamina
  - sleepiness

Where effects is how this food affects a person.

### Dessert

If a person or child eats dessert, it affects their stats as follows:

Hunger decreases by 20 points.
Sleepiness remains unchanged (0 points).
Stamina increases by 10 points.

### Meal

When a person or child eats a meal, their stats change as follows:

Hunger decreases by 50 points.
Sleepiness increases by 10 points.
Stamina increases by 25 points.

### Energizer

If a person or child consumes an energizer, their stats will adjust as follows:

Hunger decreases by 10 points.
Sleepiness decreases by 50 points.
Stamina increases by 50 points.

## An interface

Add an interface that asks the user to create a Person or a Child.

```
What did you want to create?
1. Person
2. Child
```

Once the person has been created, it will show the Person status.
I can ask them to do some activity or to eat some food.

```
My name is {{name}}, this is my status:
Hunger: {{hunger}}
Sleepyness: {{sleepyness}}
Stamina: {{stamina}}

What did you want me to do?
1. Exercise
2. Sleep
3. Study
4. Eat
5. Show status
```

If sleep, study or show status it should do what asked.
If exercise is selected you will show a list of intensities.

```
Let me know with which intensity I should exercise:
1. low
2. medium
3. hard
```

If eat is selected you will show a list of available food.

```
Which thing should I eat.
1. dessert
2. meal
3. energizer
```
