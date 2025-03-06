use rand::random;
use std::time::Instant;

const MAXN: usize = 2000;
static mut A: [[f64; MAXN]; MAXN] = [[0.0; MAXN]; MAXN];
static mut B: [f64; MAXN] = [0.0; MAXN];
static mut X: [f64; MAXN] = [0.0; MAXN];

fn initialize_inputs(n: usize) {
    unsafe {
        for i in 0..n {
            for j in 0..n {
                A[i][j] = random::<f64>();
            }
            B[i] = random::<f64>();
            X[i] = 0.0;
        }
    }
}

fn print_matrix(n: usize) {
    if n > 10 {
        return;
    }
    unsafe {
        println!("A =");
        for i in 0..n {
            for j in 0..n {
                print!("{:5.2} ", A[i][j]);
            }
            println!();
        }
        println!("\nB =");
        for i in 0..n {
            print!("{:5.2} ", B[i]);
        }
        println!();
    }
}

fn gauss(n: usize) {
    unsafe {
        for norm in 0..n - 1 {
            for row in norm + 1..n {
                let multiplier = A[row][norm] / A[norm][norm];
                for col in norm..n {
                    A[row][col] -= A[norm][col] * multiplier;
                }
                B[row] -= B[norm] * multiplier;
            }
        }

        // Substituição regressiva
        for row in (0..n).rev() {
            X[row] = B[row];
            for col in row + 1..n {
                X[row] -= A[row][col] * X[col];
            }
            X[row] /= A[row][row];
        }
    }
}

fn print_result(n: usize) {
    if n > 100 {
        return;
    }
    unsafe {
        println!("\nX =");
        for i in 0..n {
            print!("{:5.2} ", X[i]);
        }
        println!();
    }
}

fn main() {
    let n = 5; // Defina o tamanho desejado
    initialize_inputs(n);
    print_matrix(n);

    let start = Instant::now();
    gauss(n);
    let elapsed = start.elapsed();

    print_result(n);
    println!("\nElapsed time = {:.2?} ms", elapsed.as_millis());
}
