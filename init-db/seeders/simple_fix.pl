#!/usr/bin/perl
use strict;
use warnings;

my %dev2ascii = ('०'=>'0','१'=>'1','२'=>'2','३'=>'3','४'=>'4','५'=>'5','६'=>'6','७'=>'7','८'=>'8','९'=>'9');
my %ascii2dev = reverse %dev2ascii;

while (<>) {
    my $line = $_;
    
    # Rule 1: In bengla translation lines, fix: convert ascii->devanagari in VALUE only
    #  Pattern: "key": "value",
    # But first convert all devanagari to ascii everywhere (to normalize)
    foreach my $d (keys %dev2ascii) {
        $line =~ s/$d/$dev2ascii{$d}/g;
    }
    
    # Now pattern should be: "ascii_key": "ascii_value"
    # Within getBanglaTranslations section (detected by context), convert VALUE back to devanagari
    # We detect this by checking if the line looks like a translation pair
    if ($line =~ /^(\s*)"([^"]*?[0-9][^"]*)(":\s*")([^"]*?[0-9][^"]*)(".*)$/) {
        # This is a line with numbers in both key and value in translation section
        my ($indent, $key, $sep, $value, $rest) = ($1, $2, $3, $4, $5);
        
        # Convert value back to devanagari
        foreach my $a (keys %ascii2dev) {
            $value =~ s/$a/$ascii2dev{$a}/g;
        }
        
        $line = $indent . '"' . $key . '"' . $sep . $value . $rest . "\n";
    }
    
    print $line;
}
