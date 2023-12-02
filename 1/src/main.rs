use std::f32::RADIX;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let mut sum: u32 = 0;
    if let Ok(lines) = read_lines("./data.csv") {
        for line in lines {
            if let Ok(ip) = line {
                sum += get_sum(&ip);
            }
        }
    }
    println!("Sum: {}", sum);
}

fn get_sum(line: &str) -> u32 {
    let mut arr: [char; 2] = ['0', '0'];
    for ch in line.chars() {
        if ch.is_numeric() {
            if arr[0] == '0' {
                arr[0] = ch;
            }
            arr[1] = ch;
        }
    }
    return String::from_iter(arr).parse::<u32>().unwrap();
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
