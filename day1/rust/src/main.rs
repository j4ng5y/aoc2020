use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use structopt::StructOpt;
use combinations::Combinations;

fn two_number_product(n: i32, i: &[i32]) -> i32 {
    let combos: Vec<_> = Combinations::new(i.to_vec(), 2).collect();
    for combo in combos {
        let c = combo.clone().into_iter();
        let sum: i32 = c.sum();
        if sum == n {
            return &combo[0]*&combo[1]
        }
    }
    0
}

fn three_number_product(n: i32, i: &[i32]) -> i32 {
    let combos: Vec<_> = Combinations::new(i.to_vec(), 3).collect();
    for combo in combos {
        let c = combo.clone().into_iter();
        let sum: i32 = c.sum();
        if sum == n {
            return &combo[0]*&combo[1]*&combo[2]
        }
    }
    0
}

#[derive(StructOpt)]
#[structopt(name = "day1", about = "A tool to solve AOC2020 Day1")]
struct Cli {
    #[structopt(long, parse(from_os_str))]
    input: std::path::PathBuf,
    
    #[structopt(long, default_value)]
    day: i32,
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
    let mut vec: Vec<i32> = Vec::new();
    let cli = Cli::from_args();
    if let Ok(lines) = read_lines(&cli.input) {
        for line in lines {
            if let Ok(num_str) = line {
                let num = num_str.parse::<i32>().unwrap();
                vec.push(num)
            }
        }
    }

    match cli.day {
        1 => println!("Day 1 / Part 1: {}", two_number_product(2020, &vec)),
        2 => println!("Day 1 / Part 2: {}", three_number_product(2020, &vec)),
        _ => {
            println!("Day 1 / Part 1: {}", two_number_product(2020, &vec));
            println!("Day 1 / Part 2: {}", three_number_product(2020, &vec));
        }
    };
}
