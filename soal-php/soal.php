<!-- # Test PHP - Programmer [Qhomemart] #

Task :
1. Isikan proses di dalam fungsi mergeSortArray() untuk menyatukan array int a dan array int b. Lalu setelah itu di sort secara ascending.
2. Isikan proses di dalam fungsi getMissingData() untuk mencari integer yang hilang berdasarkan pola angka dari hasil fungsi mergeSortArray().
3. Isikan proses di dalam fungsi insertMissingData() untuk memasukkan integer yang hilang dari hasil fungsi getMissingData() ke dalam array hasil fungsi mergeSortArray().
4. Hasil yang diharapkan adalah pola angka yang tersusun tanpa ada integer yang hilang.

Syarat :
1. Tidak menggunakan fungsi bawaan PHP seperti 
	a. array_merge()
	b. array_push()
	c. asort()
	d. dsb.
2. Kerjakan menggunakan logic pemograman anda sendiri

Selamat Mengerjakan -->

<?php

class Test {
   
    function mergeSortArray($a, $b) {
        $result = array();
        $i = 0;
        $j = 0;

     
        while ($i < count($a) && $j < count($b)) {
            if ($a[$i] < $b[$j]) {
                $result[] = $a[$i];
                $i++;
            } else {
                $result[] = $b[$j];
                $j++;
            }
        }

    
        while ($i < count($a)) {
            $result[] = $a[$i];
            $i++;
        }

      
        while ($j < count($b)) {
            $result[] = $b[$j];
            $j++;
        }

        return $result;
    }

   
    function getMissingData($c) {
        $missing = array();
        for ($i = 0; $i < count($c) - 1; $i++) {
            $current = $c[$i];
            $next = $c[$i + 1];
          
            for ($j = $current + 1; $j < $next; $j++) {
                $missing[] = $j;
            }
        }
        return $missing;
    }

    
    function insertMissingData($c, $missing) {
        $result = array();
        $m_index = 0;
        for ($i = 0; $i < count($c); $i++) {
            $result[] = $c[$i];
            
            while ($m_index < count($missing) && $missing[$m_index] < $c[$i + 1]) {
                $result[] = $missing[$m_index];
                $m_index++;
            }
        }

        return $result;
    }

    public function main() {
        $a = array(11, 36, 65, 135, 98);
        $b = array();
        $b[0] = 81;
        $b[1] = 23;
        $b[2] = 50;
        $b[3] = 155;

       
        sort($a);
        sort($b);

       
        $c = $this->mergeSortArray($a, $b);

       
        $i = $this->getMissingData($c);

       
        $d = $this->insertMissingData($c, $i);

       
        for ($x = 0; $x < count($d); $x++) {
            echo $d[$x] . " ";
        }
    }
}

$t = new Test();
$t->main();

?>
