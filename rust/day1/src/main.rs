use std::collections::HashMap;

fn get_challenge_data(file_path: &str) -> Vec<String> {
    // Reads the file into a string
    let file_content = match std::fs::read_to_string(file_path) {
        Ok(content) => content, 
        Err(_) => { 
            eprintln!("Error reading from file. Check file name and permissions.");
            std::process::exit(2)
        }
    };

    let mut challenge_data = Vec::new();
    let lines = file_content.split("\n");
    for line in lines {
        challenge_data.push(line.to_string());
    }

    challenge_data
}

// Does part 1 of the challenge
fn part1(challenge_data: &Vec<String>) -> u32 {
    let mut sum = 0;

    for line in challenge_data {
        let mut first_digit: u32 = 10;
        let mut last_digit: u32 = 0;

        for char in line.chars() {
            if char.is_digit(10) {
                if first_digit == 10 {
                    first_digit = char.to_digit(10).unwrap();
                }
                last_digit = char.to_digit(10).unwrap();
            }
        }
        sum += (first_digit*10) + last_digit;
    }
    
    sum
}

// Does part 2 of the challenge
fn part2(challenge_data: &Vec<String>) -> u32 {
    // Mew Array with substituted values
    let mut new_lines: Vec<String> = Vec::new();

    // Value mapping - preserve the first and last char
	// e.g. 'sevenine' should be replaced with 's7nine' so the 'nine' still gets replaced
    let m = HashMap::from([
        ("one", "o1e"),
        ("two", "t2o"),
        ("three", "t3e"),
        ("four", "f4r"),
        ("five", "f5e"),
        ("six", "s6x"),
        ("seven", "s7n"),
        ("eight", "e8t"),
        ("nine", "n9e"),
    ]);

    for line in challenge_data {
        let mut new_line = line.to_string();

        for key in m.keys() {
            new_line = new_line.replace(key, m[key]);
        }

        new_lines.push(new_line);
    }

    part1(&new_lines)
}

fn main() {
    // Get the input file
    let challenge_data = get_challenge_data("input.txt");

    // Print answers
    println!("Part 1: {}", part1(&challenge_data));
    println!("Part 2: {}", part2(&challenge_data));
}
