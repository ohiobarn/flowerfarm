# Spring 2020 Plan

> Fairborn, OH is in Zone 6a.

> On average, your frost-free growing season starts Apr 11 and ends Oct 30, totaling 202 days

> Spring 2020 in Northern Hemisphere will begin on Thursday, March 19

## Flower Field Layout

Reference: 

* [dripworks planning guide](https://www.dripworks.com/resources/drip-planning-guide)
* [growing with landscape fabric](https://www.floretflowers.com/growing-with-landscape-fabric/)

![flower-field](img/field-layout.png)

* Each bed is 4 feet by 100 feet
* 10 beds total (5 beds in each irrigation zone)
* 2 feet between each bed with 6 feet between each zone 
  * Landscape fabric is 6 feet wide. This is enough to cover the beds and overlap in between to allow ~2 feet for a walking path
* 3 drip tape irrigation lines per bed
* 2 irrigation zones, each zone has 15 drip tape lines

---

## Plant Spacing

To compute the number of plants per bed; each bed is 48 inches (4 feet) wide and 1,200 inches (100 feet) long, as shown above, and by planting with 9-inch spacing, I can fit ~665 plants in each bed:

* **Plants per width**: 48in / 9in spacing = 5.3 plants; rounds down to 5 plants
* **Plants per length**: 1200in / 9in spacing = 133.3 plants; rounds down to 133 plants 
* **Plants per bed**: 5 plants X 133 plants = 665 plants per bed

>665 plants per bed, that's 6,650 plants for the entire field

![](img/plant-spacing.png)

---

## Irrigation Setup

### Flow

Knowing your flow rate is critical to sizing your system. 

I did a flow test with the spigot out in the field.  It took about 55 seconds to fill 5 gallon bucket. I used the following formula to compute gallons per hour (GPH). Or you can use the [flow calculator](https://www.dripworks.com/resources/calculators/flow-estimator) from dripworks.com

```text
55 gal   60 sec   60 min   18,000 gal
------ x ------ x ------ = ---------- ~=  327 GPH
55 sec    1 min    1 hr       55 hr

```

> Flow Capacity:  **327 GPH**

### Soil Type

Next we need to know our soil type:

* Clay - requires a 0.5 GPH emitter
* Loam - requires a 0.5-1 GPH emitter
* Sand - requires a 1-2 GPH emitter 

In my case I am not sure, I will guess loam for now.

>**Loam** - requires a 0.5-1 GPH emitter (let's go with 0.75 GPH)

### Product Selection

According to the description Drip Tape sounds best for my situation:

>**Drip Tape** is best for long, straight row crops. It is the most economical way to water your plants and is easy to install and maintain.

To start I am going to buy a kit. There seems to be two kits that could do the job:

* [drip tape row crop medium](https://www.dripworks.com/drip-tape-row-crop-kit-medium)
* [drip tape row crop large](https://www.dripworks.com/drip-tape-row-crop-kit-large)

If I only use two strips of drip tape per row then I could get the medium kit.  However, the medium kit leaves no room for error and I feel like 9 inch spacing for the plants is denser than normal so I better go with at least three strips of drip tape per row.  Given that I will get the large kit with an extra timer so I can automate two zones.

In [this blog post](https://www.floretflowers.com/soil-preparation/) Floret says they use four strips of drip tape per four foot bed but they have sandy soil and even say that you could get by with two or three strips if you don't have sandy soil, so that validates my choice.  

> I will be going with the [drip tape row crop large kit](https://www.dripworks.com/drip-tape-row-crop-kit-large)

### Zones

The drip tape that comes with the [large kit](https://www.dripworks.com/drip-tape-row-crop-kit-large) uses [drip tape part# TA1508L](https://www.dripworks.com/aqua-traxx-15-mil-drip-tape-rolls) which is rated at **Flow per 100 feet is 20 GPH**. So that is 20 GPH for each strip of drip tape. I need 30 strips total for the field which would require 600 GPH (20 GPH X 30 strips), but my system can only generate 327 GPH. Therefore, if I split my field into two zone, 5 beds requiring 15 strips each, then each zone will only need 300 GPH (20 GPH X 15 strips) which is less than 327 GPH, so we are good.

> two zones, each requiring 300 GPH (20 GPH X 15 strips)

---

## Germination Station

### One Shelf

The image below shows the layout for one shelf of the germination station. Each shelf will be able to hold `288 plants`, four starter trays of 72 plants each.  

> 288 plants per shelf

![](img/one-shelf-seed-starter.png)


### One Station

Each station is a five shelf unit.  The top shelf is used for storage and four shelves will be used for germination.  Each station will be able to hold `1,152 plants`, four shelves of 288 plants each.

> 1,152 plants per station

---

## Beds

* main cutting beds
* pumpkin patch
* sunflower strip

## Plants

Main cutting beds, annuals with irrigation.

> * [Main field layout PDF](Main-Field-2020.pdf)
> * [Plant list PDF](Main-Field-2020-Plant-List.pdf)


Plant Spacing:

| #| Plant               | Propagation | Start   | Between plants in row   | Between rows | Rows per bed | Fabric | Irrigation |
|--|---------------------|-------------|---------|-------------------------|--------------|--------------|--------|------------|
| 1| Dahlias             | tuber       | outside | 16 in                   | 16 in        | 3            |  yes   |   yes      |
| 2| Dahlia Mix          | seed        | inside  | 12 in                   | 16 in        | 3            |  yes   |   yes      |
| 3| Ranunculus          | corms       | outside |  8 in                   | 10 in        | 4            |   no   |   yes      |  
| 4| Anemone             | corms       | outside |  6 in                   | 10 in        | 4            |   no   |   yes      |
| 5| Zinnias             | seed        | inside  |  9 in                   |  9 in        | 5            |  yes   |   yes      |
| 6| Black Eyed Susans   | seed        | inside  |  9 in                   |  9 in        | 5            |  yes   |   yes      |
| 7| Snapdragon          | seed        | inside  |  9 in                   |  9 in        | 5            |  yes   |   yes      |
| 8| Cosmos              | | |
| 9| Stocks              | | |
|10| Celosia             | | |
|11| Gormphrena          | | |
|12| Larkspur            | | |
|13| Marigold            | | |
|14| Choc.Qn Anne's Lace | | |
|15| Sweet Pea           | | |

Varieties:

In this table if the number of plants is fixed then the *bed-linear-feed* will be computed in the Garden Planner app.  However in most cases the *bed-linear-feed* is given and the number of plans is computed. **Note** a "plant" might require several seeds (or tubers or corms) if it is a "thinning" plant like the pumpkins. 


| Plant                  | Variety                     | Fixed Qty   | Row Group |
|------------------------|-----------------------------|-------------|-----------|
| Dahlias                | Jan Ryecroft (Sm 4”)        | 5 tubers    |    6A
|                        | Lady Natalie (Sm 4”)        | 5 tubers    |    6A
|                        | Golden Scepter (Sm 2”)      | 5 tubers    |    6A
|                        | Black Satin (Sm 2-4”)       | 5 tubers    |    6B
|                        | Cornel (Sm 4”)              | 5 tubers    |    6B
|                        | Dorothy (Sm)                | 5 tubers    |    6B
|                        | Franz Kafka (Sm 2-3”)       | 7 tubers    |    6C
|                        | Wizard of Oz (Sm 3”)        | 3 tubers    |    6C
|                        | Peaches n Cream (Med 4-6“)  | 8 tubers    |    6D
|                        | Conn.Coral (Med)            | 8 tubers    |    6D
|                        | Mystique (Med 4-6“)         | 8 tubers    |    6D
|                        | BoomBoom White (Med 6”)     | 5 tubers    |    6E
|                        | Sweet Nathalie (Med. 5”)    | 5 tubers    |    6E
|                        | L’ancresse (Med 4-6“)       | 5 tubers    |    6E
|                        | Thomas Edison (Med 6”)      | 5 tubers    |    6F
|                        | Shiloh Noelle (XL 10”)      | 5 tubers    |    6F
|                        | Hamari Gold (XL 10”)        | 5 tubers    |    6F
|                        | Penhill Watermelon (XL 10”) | 8 tubers    |    6G
|                        | Otto’s Thrill (XL)          | 8 tubers    |    6G
|                        | Café Au Lait (XL)           | 8 tubers    |    6G
|                        | Mix (seed)                  | 90 seeds    |    6H


TODO - use [this info for the dahlia seed and work into the tables](https://www.johnnyseeds.com/flowers/dahlia/giant-hybrid-mix-dahlia-seed-3397.11.html)

---

## Filler

Floret said 50% Foliage and filler

| #| Plant          | | |
|--|----------------|-|-|
| 1| Baby's Breath (Gypsophila) | | |
| 2| Bishops Flower | | |
| 3| Love-in-a-Mist | | |

## Foliage

| #| Plant            | | |
|--|------------------|-|-|
| 1| Dusty Miller     | | perennials but grown as annual|
| 2| Scented Geranium | | |
| 3| Grasses-Panicum  | | "frosted explosion" |
| 4| Eucalyptus       | | "silver drop", "seeded" |

## Sunflowers

Sunflower beds, no irrigation

| #| Variety             | | |
|--|------------------|-|-|
| 1| Single bloom | | |
| 2| Blush | | |

## Pumpkin Patch

Pumpkin Patch, no irrigation

|#| Variety        | Size   | Days to Maturity | Vine Length | Between plants in row | Between rows | Grow    | Link  |
|-|----------------|--------|------------------|-------------|-----------------------|--------------|---------|-------|
|1| PipSqueak      | Small  | 100              | Medium      | 33 in                 |  8 ft        | &#8224; | [JS](https://www.johnnyseeds.com/vegetables/pumpkins/pipsqueak-pmr-f1-pumpkin-seed-4019.html?cgid=pumpkins)
|2| Racer Treated  | Medium |  85              | Short       | 33 in                 |  6 ft        | &#8224; | [JS](https://www.johnnyseeds.com/vegetables/pumpkins/jack-o-lanterns/racer-treated-f1-pumpkin-seed-2163T.html#q=racer&lang=en_US&start=1)
|3| Howden         | Large  | 115              | Long        | 60 in                 | 12 ft        | &#8224; | [JS](https://www.johnnyseeds.com/vegetables/pumpkins/specialty-pumpkins/casperita-f1-pumpkin-seed-3076.html#q=casperita&lang=en_US&start=1)
|4| Casperita      | Mini   |  77              | Medium      | 18 in                 | 10 ft        | &#8224; | [JS](https://www.johnnyseeds.com/vegetables/pumpkins/specialty-pumpkins/casperita-f1-pumpkin-seed-3076.html#q=casperita&lang=en_US&start=1)
|5| Flat Stacker   | Medium |  95              | Long        | 42 in                 | 12 ft        | &#8224; | [JS](https://www.johnnyseeds.com/vegetables/pumpkins/specialty-pumpkins/flat-stacker-f1-pumpkin-seed-3574.html)
|6| Ornamental Mix | Gourd  |  95              | Medium      | 60 in                 |  8 ft        | &#8224; | [JS](https://www.johnnyseeds.com/vegetables/gourds/ornamental-gourds-small-mixed-gourd-seed-604.html#q=ornamental%2Bgourds&lang=en_US&start=1)

> &#8224; - For direct seeding, sow 2 seeds per spacing, 1/2-1" deep. Thin to 1 plant per spacing after seedlings are established

### Pumpkin Grow Sow Chart

![](img/Pumpkin_Patch_Plant_List.png )

### Pumpkin Grow Sow Field Layout

![](img/Pumpkin_Patch_Field.png)
---

![](img/pumpkins-specialty-varieties-comparison-chart_pdf.png)


## Resources

* [Floret-Grow-Great-Dahlias.pdf](Floret-Grow-Great-Dahlias.pdf)
